#!/usr/bin/env bash

set -euo pipefail

mapdir=`mktemp -d`
database_url=${DATABASE_URL}
postgres_user=`echo $database_url | awk -F ':' '{print $2}' | cut -b 3-`
postgres_password=`echo $database_url | awk -F ':' '{print $3}' | awk -F '@' '{print $1}'`
postgres_db=`echo $database_url | awk -F '/' '{print $4}' | awk -F '?' '{print $1}'`
export PGPASSWORD=${postgres_password}
psql="psql -U ${postgres_user} ${postgres_db}"
migrations_dir=${MIGRATIONS_DIR:="/migrations"}

alias clickhouse="docker run -i --rm --link clickhouse:clickhouse curlimages/curl 'http://clickhouse:8123/?query=' -s --data-binary @-"

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
        $psql --command "BEGIN; ${UP}; COMMIT;" && echo "Migration applied."
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

    case "${option}" in
        up|UP)
            UP=$(sed '/-- <up>/,/-- <\/up>/!d' ${filename} | grep -v '^--')
            echo -e "\nSQL:\n${UP}\n"
            read -p "Confirm? y|n: " confirm
            if [ "${confirm}" == "y" ]; then
                $psql -c "BEGIN; ${UP}; COMMIT;" && echo "Migration applied."
            fi
            ;;
        down|DOWN)
            DOWN=$(sed '/-- <down>/,/-- <\/down>/!d' ${filename} | grep -v '^--')
            echo -e "\nSQL:\n${DOWN}\n"
            read -p "Confirm? y|n: " confirm
            if [ "${confirm}" == "y" ]; then
                $psql -c "BEGIN; ${DOWN}; COMMIT;" && echo "Migration reverted."
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
    shell) $psql -q -X -v VERBOSITY=terse -v ON_ERROR_STOP=1 -A -t -w ;;
    help) helpme ;;
    *) show_menu ;;
esac
