/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose-zap
 * @Date:        2023-07-26 09:31
 * @Description:
 */

package rzap

type FileConfig struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	LocalTime  bool
	Compress   bool
}

type Option func(f *FileConfig)

func WithFileName(filename string) Option {
	return func(f *FileConfig) {
		if len(filename) > 0 {
			f.FileName = filename
		}
	}
}

func WithMaxSize(maxSize int) Option {
	return func(f *FileConfig) {
		f.MaxSize = maxSize
	}
}

func WithMaxBackups(maxBackups int) Option {
	return func(f *FileConfig) {
		f.MaxBackups = maxBackups
	}
}

func WithMaxAge(maxAge int) Option {
	return func(f *FileConfig) {
		f.MaxAge = maxAge
	}
}

func WithLocalTime(enabled bool) Option {
	return func(f *FileConfig) {
		f.LocalTime = enabled
	}
}

func WithCompress(enabled bool) Option {
	return func(f *FileConfig) {
		f.Compress = enabled
	}
}
