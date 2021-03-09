package tmvalerrs

import (
	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/misc/dict"
	"github.com/NoHomey/chaos-go-camp-proj/misc/validator/valerrs"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Ctx is the data needed for LogAndReturnCtxErr.
type Ctx struct {
	Err    validator.ValidationErrors
	Logger *zap.Logger
	Msg    string
	Log    []zapcore.Field
}

//LogAndReturnCtxErr logs the error and transforms it in ctxerror.Error.
func LogAndReturnCtxErr(ctx *Ctx) ctxerr.Error {
	invalid := valerrs.Collect(ctx.Err)
	fields, validators := dict.Data(invalid)
	k := len(ctx.Log)
	log := make([]zapcore.Field, k, k+3)
	copy(log, ctx.Log)
	log = append(
		log,
		zap.Strings("invalidFields", fields),
		zap.Strings("failedValidators", validators),
		zap.Error(ctx.Err),
	)
	ctx.Logger.Error(ctx.Msg, log...)
	return ctxerr.NewInvalData(ctx.Err, invalid)
}
