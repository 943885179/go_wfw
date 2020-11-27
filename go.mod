module qshapi

go 1.15

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/Unknwon/goconfig v0.0.0-20200908083735-df7de6a44db8
	github.com/bitly/go-simplejson v0.5.0
	github.com/boombuler/barcode v1.0.1-0.20190219062509-6c824513bacc
	github.com/denisenkom/go-mssqldb v0.0.0-20200910202707-1e08a3fab204 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/disintegration/imaging v1.6.2
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/gin-contrib/gzip v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.3
	github.com/goki/freetype v0.0.0-20181231101311-fa8a33aabaff
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20190910122728-9d188e94fb99 // indirect
	github.com/jackc/pgx v3.6.0+incompatible // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jordan-wright/email v4.0.1-0.20200917010138-e1c00e156980+incompatible
	github.com/kr/pretty v0.2.0 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.4 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.16.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/mojocn/base64Captcha v1.3.1
	github.com/mozillazg/go-pinyin v0.18.0
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/shopspring/decimal v0.0.0-20200227202807-02e2044944cc // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/streadway/amqp v1.0.0 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/tebeka/strftime v0.1.5 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go v1.0.20
	github.com/tuotoo/qrcode v0.0.0-20190222102259-ac9c44189bf2
	github.com/typa01/go-utils v0.0.0-20181126045345-a86b05b01c1e
	github.com/vladoatanasov/logrus_amqp v0.0.0-20181023103017-b21faf6f8ae3
	github.com/willf/bitset v1.1.11 // indirect
	github.com/yanyiwu/gojieba v1.1.2
	go.etcd.io/bbolt v1.3.5 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee // indirect
	golang.org/x/image v0.0.0-20200927104501-e162460cd6b5
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20201118182958-a01c418693c7 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.32.0
	gopkg.in/olivere/elastic.v5 v5.0.84
	gopkg.in/sohlich/elogrus.v2 v2.0.2
	gorm.io/driver/mysql v1.0.2
	gorm.io/gorm v1.20.2
)

// 替换为v1.26.0版本的gRPC库
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
