package mysql

import (
	"cloudmer/share"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// mysql 驱动
type MysqlDriver struct {
	// host
	Host string
	// port
	Port int
	// username
	Username string
	// password
	Password string
	// dbname
	Dbname string
	// charset
	Charset string
	// 连接timeout时间
	Timeout string
	// 读取超时时间
	ReadTimeout string
	// 表前缀
	TablePrefix string
	// 设置空闲连接池中连接的最大数量
	MaxIdle int
	// 设置打开数据库连接的最大数量
	MaxOpen int
	// 设置连接可复用的最大时间 单位分钟
	MaxLifetime int
	// 日志级别 Silent=1 Error=2 Warn=3 Info=4
	LogLevel int
}

func (driver *MysqlDriver) Build() *gorm.DB {
	if err := driver.validate(); err != nil {
		panic(err)
	}
	// dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s&readTimeout=%s",
		driver.Username,
		driver.Password,
		driver.Host,
		driver.Port,
		driver.Dbname,
		driver.Charset,
		driver.Timeout,
		driver.ReadTimeout,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		// 打开sql日志
		Logger: logger.Default.LogMode(logger.LogLevel(driver.LogLevel)),
		// 表前缀
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: driver.TablePrefix,   // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(driver.MaxIdle)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(driver.MaxOpen)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。 单位分钟
	sqlDb.SetConnMaxLifetime(time.Duration(driver.MaxLifetime) * time.Minute)

	err = sqlDb.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

// 参数验证
func (driver *MysqlDriver) validate() error {
	if driver.Host == "" {
		return errors.New("mysql.host 未设置")
	}
	if driver.Port == 0 {
		return errors.New("mysql.port 未设置")
	}
	if driver.Username == "" {
		return errors.New("mysql.username 未设置")
	}
	if driver.Password == "" {
		return errors.New("mysql.password 未设置")
	}
	if driver.Dbname == "" {
		return errors.New("mysql.dbname 未设置")
	}
	if driver.Charset == "" {
		return errors.New("mysql.charset 未设置")
	}
	if driver.LogLevel == 0 {
		return errors.New("mysql.logLevel 未设置")
	}
	if driver.LogLevel >= 5 || driver .LogLevel < 0 {
		return errors.New("mysql.logLevel not in [1,2,3,4]")
	}
	if driver.MaxIdle == 0 {
		return errors.New("mysql.maxIdle 未设置")
	}
	if driver.MaxOpen == 0 {
		return errors.New("mysql.maxOpen 未设置")
	}
	if driver.MaxLifetime == 0 {
		return errors.New("mysql.maxLifetime 未设置")
	}
	return nil
}

func StdConfig(key string) *MysqlDriver {
	mysqlDriver := DefaultConfig()
	if err := share.Viper.UnmarshalKey(key, mysqlDriver); err != nil {
		panic(err)
	}
	return mysqlDriver
}

func DefaultConfig() *MysqlDriver {
	return &MysqlDriver{
		// 字符类型
		Charset: "utf8",
		// 设置空闲连接池中连接的最大数量
		MaxIdle: 3,
		// 设置打开数据库连接的最大数量
		MaxOpen: 3,
		// 设置连接可复用的最大时间 单位分钟
		MaxLifetime: 3,
		// 日志级别 Silent=1 Error=2 Warn=3 Info=4
		LogLevel: 4,
		// 连接timeout时间
		Timeout: "5s",
		// 读取超时时间
		ReadTimeout: "60s",
	}
}