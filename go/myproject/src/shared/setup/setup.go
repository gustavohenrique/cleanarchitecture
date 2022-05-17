package setup

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"myproject/src/shared/cryptus"
	"myproject/src/shared/datetime"
	"myproject/src/shared/uuid"
)

func CreateSqliteDB(file, schema string) error {
	conn, err := getSqliteConnection(file)
	if err != nil {
		return err
	}
	_, err = conn.Exec(schema)
	return err
}

func InsertAdminInSqliteDB(file, email, password string) error {
	conn, err := getSqliteConnection(file)
	if err != nil {
		return err
	}
	name := "Super Admin"
	roleID := uuid.NewV4()
	permissions := `[{"action": "admin", "target": "*"}]`
	q := "INSERT INTO roles (id, name, permissions) VALUES (?, ?, ?)"
	_, err = conn.Exec(q, roleID, name, permissions)
	if err != nil {
		return err
	}

	groupID := uuid.NewV4()
	roles := fmt.Sprintf(`["%s"]`, roleID)
	q = "INSERT INTO groups (id, name, roles_ids) VALUES (?, ?, ?)"
	_, err = conn.Exec(q, groupID, name, roles)
	if err != nil {
		return err
	}

	userID := uuid.NewV4()
	priv, pub, err := cryptus.GenerateRsaKeyPair(2048)
	if err != nil {
		return err
	}
	privateKey := cryptus.FromBytesToBase64(priv)
	publicKey := cryptus.FromBytesToBase64(pub)
	createdAt := datetime.ToString(datetime.Now())
	hashBytes, saltBytes := cryptus.HashPbkdf2(password+createdAt, []byte(""))
	salt := cryptus.ToHex(saltBytes)
	hash := cryptus.ToHex(hashBytes)
	pass := hash + salt
	groups := fmt.Sprintf(`["%s"]`, groupID)
	q = `INSERT INTO users (
		id,
		created_at,
		email,
		password,
		full_name,
		groups_ids,
		salt,
		private_key,
		public_key,
		last_login_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)`
	_, err = conn.Exec(q, userID, createdAt, email, pass, name, groups, salt, privateKey, publicKey)
	return err
}

func getSqliteConnection(file string) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("sqlite3", file)
	return conn, err
}
