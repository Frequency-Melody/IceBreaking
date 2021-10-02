package main

import "IceBreaking/router"

// @title 认人脸后端`
// @version 1.0`
// @description 认人脸后端`
// @termsOfService [http://swagger.io/terms/](http://swagger.io/terms/)`
// @contact.name BirdBirdLee`
// @contact.url [http://www.swagger.io/support](http://www.swagger.io/support)`
// @contact.email li_hepeng@qq.com`
// @license.name Apache 2.0`
// @license.url [http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)`
// @host http://api.aflybird.cn:8090`
// @BasePath /ice`
func main() {
	r := router.Router{}
	r.Run(8091)
}
