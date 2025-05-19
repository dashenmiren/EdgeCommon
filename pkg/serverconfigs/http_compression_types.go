package serverconfigs

type HTTPCompressionType = string

const (
	HTTPCompressionTypeGzip    HTTPCompressionType = "gzip"
	HTTPCompressionTypeDeflate HTTPCompressionType = "deflate"
	HTTPCompressionTypeBrotli  HTTPCompressionType = "brotli"
	HTTPCompressionTypeZSTD    HTTPCompressionType = "zstd"
)
