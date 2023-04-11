package errors

import (
	"context"
	"github.com/Yostardev/errors/store/ent"
	"github.com/Yostardev/errors/store/ent/errorcode"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
)

var (
	db *ent.Client
)

func initDB(dsn, logLevel string) {
	var err error
	db, err = ent.Open("mysql", dsn, ent.Log(func(data ...any) {
		zapLogger.Debug(data)
	}))
	if err != nil {
		zapLogger.Panicf("failed opening connection to mysql: %v", err)
	}

	if logLevel == "debug" {
		db = db.Debug()
	}

	if err = db.Schema.Create(context.Background()); err != nil {
		zapLogger.Panicf("failed creating schema resources: %v", err)
	}

	ctx := context.Background()

	count, err := db.ErrorCode.Query().Where(errorcode.Name("Err_Unknown")).Count(ctx)
	if err != nil {
		zapLogger.Panicf("failed get resources: %v", err)
	}

	if count == 0 {
		err = Insert(ctx, 50000, 2, "Err_Unknown", "未知错误")
		if err != nil {
			zapLogger.Panicf("failed creating resources: %v", err)
		}
	}

	count, err = db.ErrorCode.Query().Where(errorcode.Name("Err_GRPC_Connection")).Count(ctx)
	if err != nil {
		zapLogger.Panicf("failed get resources: %v", err)
	}

	if count == 0 {
		err = Insert(ctx, 50001, 14, "Err_GRPC_Connection", "无法连接至Grpc服务")
		if err != nil {
			zapLogger.Panicf("failed creating resources: %v", err)
		}
	}
}

func closeDB() {
	if err := db.Close(); err != nil {
		zapLogger.Warn(err)
	}
}

func Insert(ctx context.Context, errorCode int, grpcCode codes.Code, name, message string) error {
	if err := db.ErrorCode.Create().
		SetName(name).
		SetErrorCode(errorCode).
		SetGrpcStatus(uint32(grpcCode)).
		SetMessage(message).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}

func Update(ctx context.Context, id int, errorCode int, grpcCode codes.Code, name, message string) error {
	if err := db.ErrorCode.Update().
		Where(errorcode.ID(id)).
		SetName(name).
		SetErrorCode(errorCode).
		SetGrpcStatus(uint32(grpcCode)).
		SetMessage(message).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}

func Remove(ctx context.Context, id int) error {
	if _, err := db.ErrorCode.Delete().Where(errorcode.ID(id)).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func FetchAll(ctx context.Context) ([]*ent.ErrorCode, error) {
	return db.ErrorCode.Query().All(ctx)
}
