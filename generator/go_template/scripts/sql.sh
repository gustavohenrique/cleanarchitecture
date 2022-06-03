#!/usr/bin/env bash

set -euo pipefail

db_engine=${DB_ENGINE:-"sqlite"}
database_url=`echo ${DATABASE_URL} | awk -F '?' '{print $1}'`
mapdir=`mktemp -d`
migrations_dir=${MIGRATIONS_DIR:=${PWD}}

function put {
    [ "$#" != 3 ] && exit 1
    mapname=$1; key=$2; value=$3
    [ -d "${mapdir}/${mapname}" ] || mkdir "${mapdir}/${mapname}"
    echo $value >"${mapdir}/${mapname}/${key}"
}

function get {
    [ "$#" != 2 ] && exit 1
    mapname=$1; key=$2
    cat "${mapdir}/${mapname}/${key}"
}

function is_empty {
    local dir="${1}"
    shopt -s nullglob
    local files=( "${dir}"/* "${dir}"/.* )
    [[ ${#files[@]} -eq 2 ]]
}

function apply_all {
    if [ ! -d ${migrations_dir} ]; then
        echo "[MIGRATIONS_DIR=${migrations_dir}]. Directory not found."
        exit 0
    fi
    if is_empty ${migrations_dir}; then
        echo "[MIGRATIONS_DIR=${migrations_dir}]. Directory is empty."
        exit 0
    fi
    for f in `/bin/ls -1 ${migrations_dir}/[0-9]* | sort`; do
        name=`basename ${f}`
        echo -e "${name}"
        UP=$(sed '/-- <up>/,/-- <\/up>/!d' ${f} | grep -v '^--')
        if [ "${db_engine}" == "postgres" ]; then
          psql ${database_url} -c "BEGIN; ${UP}; COMMIT;" && echo "Migration applied."
        else
          sqlite3 ${database_url} "BEGIN; ${UP}; COMMIT;" && echo "Migration applied."
        fi
    done
}

function show_menu {
    if [ ! -d ${migrations_dir} ]; then
        echo "[MIGRATIONS_DIR=${migrations_dir}]. Directory not found."
        exit 0
    fi
    if is_empty ${migrations_dir}; then
        echo "[MIGRATIONS_DIR=${migrations_dir}]. Directory is empty."
        exit 0
    fi
    echo -e "ID\tFilename"
    i=1
    for f in `/bin/ls -1 ${migrations_dir}/[0-9]* | sort`; do
        name=`basename ${f}`
        echo -e "${i}\t${name}"
        put "files" "${i}" "${f}"
        ((i=i+1))
    done
    echo
    read -p "ID: " key
    filename=$(get "files" "${key}")
    echo "${filename}"
    # source ${filename}

    read -p "Apply (up) or rollback (down)? up|down: " option

    if [ -z "${database_url}" ]; then
        read -p "DATABASE_URL: " db_url
        database_url=`echo ${db_url} | awk -F '?' '{print $1}'`
    fi

    case "${option}" in
        up|UP)
            UP=$(sed '/-- <up>/,/-- <\/up>/!d' ${filename} | grep -v '^--')
            echo -e "\nSQL:\n${UP}\n"
            read -p "Confirm? y|n: " confirm
            if [ "${confirm}" == "y" ]; then
                if [ "${db_engine}" == "postgres" ]; then
                    psql ${database_url} -c "BEGIN; ${UP}; COMMIT;" && echo "Migration applied."
                else
                    sqlite3 ${database_url} "BEGIN; ${UP}; COMMIT;" && echo "Migration applied."
                fi
            fi
            ;;
        down|DOWN)
            DOWN=$(sed '/-- <down>/,/-- <\/down>/!d' ${filename} | grep -v '^--')
            echo -e "\nSQL:\n${DOWN}\n"
            read -p "Confirm? y|n: " confirm
            if [ "${confirm}" == "y" ]; then
                if [ "${db_engine}" == "postgres" ]; then
                    psql ${database_url} -c "BEGIN; ${DOWN}; COMMIT;" && echo "Migration reverted."
                else
                    sqlite3 ${database_url} "BEGIN; ${DOWN}; COMMIT;" && echo "Migration reverted."
                fi
            fi
            ;;
        *) echo "Invalid option."; exit 0 ;;
    esac
}

function helpme {
    echo "HELP"
    echo "-----"
    echo "Run the command bellow to see the migrations available:"
    echo "migrate ls"
}

case "${1}" in
    menu) show_menu ; exit 0 ;;
    show|ls|view) /bin/ls -1 [0-9]* | sort; exit 0 ;;
    apply) apply_all ; exit 0 ;;
    help) helpme ;;
    *) show_menu ;;
esac
