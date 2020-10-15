module qshapi

go 1.15

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/Unknwon/goconfig v0.0.0-20200908083735-df7de6a44db8
	github.com/bitly/go-simplejson v0.5.0
	github.com/blevesearch/bleve v1.0.10 // indirect
	github.com/boombuler/barcode v1.0.1-0.20190219062509-6c824513bacc
	github.com/chai2010/webp v1.1.0 // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20200910202707-1e08a3fab204 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/disintegration/imaging v1.6.2
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/gin-contrib/gzip v0.0.3
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.3
	github.com/gofiber/fiber v1.14.6 // indirect
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/goki/freetype v0.0.0-20181231101311-fa8a33aabaff
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/jackc/pgx v3.6.0+incompatible // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jordan-wright/email v4.0.1-0.20200917010138-e1c00e156980+incompatible
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mattn/go-sqlite3 v1.14.4 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.16.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/grpc-go v0.0.0-20190130160115-549af9fb4bf2 // indirect
	github.com/micro/protoc-gen-micro v1.0.0 // indirect
	github.com/mojocn/base64Captcha v1.3.1
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.7.0
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/streadway/amqp v1.0.0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/tebeka/strftime v0.1.5 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go v1.0.20
	github.com/tuotoo/qrcode v0.0.0-20190222102259-ac9c44189bf2
	github.com/typa01/go-utils v0.0.0-20181126045345-a86b05b01c1e
	github.com/ugorji/go v1.1.7 // indirect
	github.com/vladoatanasov/logrus_amqp v0.0.0-20181023103017-b21faf6f8ae3
	github.com/wangbin/jiebago v0.3.2 // indirect
	github.com/willf/bitset v1.1.11 // indirect
	github.com/xyproto/permissions2 v0.0.0-20200902135438-05029d08c3f2
	github.com/yanyiwu/gojieba v1.1.2
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee // indirect
	golang.org/x/image v0.0.0-20200927104501-e162460cd6b5
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.23.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/olivere/elastic.v5 v5.0.84
	gopkg.in/sohlich/elogrus.v2 v2.0.2
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gorm.io/driver/clickhouse v0.0.0-20201012085455-facfac3584cc
	gorm.io/driver/mysql v1.0.2
	gorm.io/driver/postgres v1.0.2
	gorm.io/driver/sqlite v1.1.3
	gorm.io/driver/sqlserver v1.0.4
	gorm.io/gorm v1.20.2
	gorm.io/hints v0.0.0-20201009065012-5a8ac6261297 // indirect
	gorm.io/plugin/dbresolver v1.0.0
)

// 替换为v1.26.0版本的gRPC库
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
