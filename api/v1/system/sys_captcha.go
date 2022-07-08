package system

import (
	"bifrost/server/global"
	"bifrost/server/model/common/response"
	systemRes "bifrost/server/model/system/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

type Student struct {
	Username string `bson:"username"`
	Tier     int    `bson:"tier"`
}

// Captcha
// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=systemRes.SysCaptchaResponse,msg=string} "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router /base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysCaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.GVA_CONFIG.Captcha.KeyLong,
		}, "验证码获取成功", c)
	}
}

// test api
func (b *BaseApi) Test(c *gin.Context) {

	//打印日志
	//global.GVA_LOG.Error("test info log")

	/*c.JSON(http.StatusOK, gin.H{
		//Message String `json:"ok"`
		"message": "ok1",
	})*/

	//(mysql)查询user
	//u := &system.SysUser{Username: "admin", Password: "123456"}
	//var user system.SysUser
	//global.GVA_DB.Where("username = ?", u.Username).First(&user)

	//连接mongodb
	/*opts := options.Client().ApplyURI("mongodb://120.78.144.133:27019")
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}
	//选择集合
	dbName := "test"
	dbNow := client.Database(dbName)*/
	//查看所有集合
	/*AllCollections, _ := dbNow.ListCollectionNames(context.Background(), bson.D{}) //此处可以获得mongodb中数据库中所有集合的名字，以列表的形式返回
	for _, collection := range AllCollections {
		fmt.Println(collection)
	}*/

	//连接库
	//client := utils.Connection()

	//设置表
	//usersTable := utils.Table(client, "users")
	//查询
	//usersTable.FindOne()
	/*	var result Student
		filter := bson.M{"name": "n1"}
		err := usersTable.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Found multiple documents (array of pointers): %#v\n", result)*/

	//filter := bson.M{"tier": 100}
	//fmt.Printf("Found a single document: %+v\n", result)

	/*for cursor.Next(context.TODO()) {
		var elem Student
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	cursor.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %#v\n", results)*/
	//处理结果
	//var meta []metadata
	//err = cursor.All(ctx, &meta)
	/*	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})*/

	/*ResultsMap := make(map[string]interface{})
	for _, collection := range AllCollections {
		ConnectSet := dbNow.Collection("users") //连接数据库里的集合
		cursor, _ := ConnectSet.Find(context.Background(), filter)
		defer cursor.Close(context.Background())

		//创建bson.M类型的数组
		var temp []bson.M

		if err = cursor.All(context.Background(), &temp); err != nil {
			log.Fatal(err)
		}

		if len(temp) > 0 {
			//一个表单创建一个map
			OneConnectSet := make([]interface{}, len(temp))

			//做相应的处理即可

			ResultsMap[collection] = OneConnectSet
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": &ResultsMap,
	})*/
	/*	opts := options.Client().ApplyURI("mongodb://localhost:27019")
		//连接
		client, err := mongo.Connect(context.Background(), opts)
		if err != nil {
			response.FailWithMessage("mongo连接失败", c)
		}

		fmt.Println(client.Ping(context.Background(), readpref.Primary()))
	*/
	//dbName := "test"

}
