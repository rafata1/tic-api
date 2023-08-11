package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/url"
	"strings"
)

// MySQLOption for MySQL options
type MySQLOption struct {
	Key   string `mapstructure:"key"`
	Value string `mapstructure:"value"`
}

// MySQLConfig for configuring MySQL
type MySQLConfig struct {
	Host         string        `mapstructure:"host"`
	Port         uint16        `mapstructure:"port"`
	Database     string        `mapstructure:"database"`
	Username     string        `mapstructure:"username"`
	Password     string        `mapstructure:"password"`
	MaxOpenConns int           `mapstructure:"max_open_conns"`
	MaxIdleConns int           `mapstructure:"max_idle_conns"`
	Options      []MySQLOption `mapstructure:"options"`
}

// DSN returns data source name
func (c MySQLConfig) DSN() string {
	var opts []string
	for _, o := range c.Options {
		key := url.QueryEscape(o.Key)
		value := url.QueryEscape(o.Value)
		opts = append(opts, key+"="+value)
	}
	optStr := strings.Join(opts, "&")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Database, optStr)
}

// MustConnect connects to database using sqlx
func (c MySQLConfig) MustConnect() *sqlx.DB {
	db := sqlx.MustOpen("mysql", c.DSN())

	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)
	return db
}
