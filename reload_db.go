package zeroweb

import (
	"github.com/jackc/pgx"
	"github.com/rs/zerolog/log"
)

func (a *Zeroweb) ReloadDB() {
	var config pgx.ConnPoolConfig
	config.Host = a.Config.GetString("db.host")
	config.User = a.Config.GetString("db.user")
	config.Password = a.Config.GetString("db.password")
	config.Database = a.Config.GetString("db.database")
	config.Port = uint16(a.Config.GetInt64("db.port"))
	config.MaxConnections = a.Config.GetInt("db.maxconnections")

	if a.DBConfig != nil &&
		config.Host == a.DBConfig.Host &&
		config.User == a.DBConfig.User &&
		config.Password == a.DBConfig.Password &&
		config.Database == a.DBConfig.Database &&
		config.Port == a.DBConfig.Port &&
		config.MaxConnections == a.DBConfig.MaxConnections {
		return // config didn't change for DB
	}

	//TODO
	// config.AfterConnect = func(conn *pgx.Conn) error {
	// 	worldSelectStmt = mustPrepare(conn, "worldSelectStmt", "SELECT id, randomNumber FROM World WHERE id = $1")
	// 	worldUpdateStmt = mustPrepare(conn, "worldUpdateStmt", "UPDATE World SET randomNumber = $1 WHERE id = $2")
	// 	fortuneSelectStmt = mustPrepare(conn, "fortuneSelectStmt", "SELECT id, message FROM Fortune")
	// 	return nil
	// }
	//TODO
	// config.DialFunc

	connPool, err := pgx.NewConnPool(config)
	if err != nil {
		if a.DB == nil {
			log.Fatal().Err(err).Interface("config", config).Msg("connection to DB failed")
		}
		log.Error().Err(err).Interface("config", config).Msg("connection to DB failed (keeping old db connection)")
		return
	}
	a.DBConfig = &config
	a.DB = connPool
}
