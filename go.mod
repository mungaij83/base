module github.com/hexya-addons/base

go 1.18

require (
	github.com/google/uuid v1.1.1
	github.com/hexya-addons/web v0.1.7
	github.com/hexya-erp/hexya v0.1.7
	github.com/hexya-erp/pool v1.0.2
	github.com/smartystreets/goconvey v0.0.0-20190306220146-200a235640ff
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.5.0
)


replace github.com/hexya-erp/hexya v0.1.7 => ../../hexya/
replace github.com/hexya-addons/web v0.1.7 => ../../hexya-addons/web/