/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose-zap
 * @Date:        2023-07-25 17:32
 * @Description:
 */

package rzap

import "go.uber.org/zap/zapcore"

func GetEncode(isJson bool, cfg zapcore.EncoderConfig) zapcore.Encoder {

	//cfg.EncodeName

	if isJson {
		return zapcore.NewJSONEncoder(cfg)
	} else {
		return zapcore.NewConsoleEncoder(cfg)
	}
}
