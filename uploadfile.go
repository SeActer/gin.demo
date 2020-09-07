package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		//c.SaveUploadedFile(file,"./" + file.Filename)
		in, _ := file.Open()
		out, _ := os.Create("./" + file.Filename)
		io.Copy(out, in)
		defer in.Close()
		defer out.Close()
		c.Writer.Header().Add("Content-Disposition", fmt.Sprint("attachment;filename=%s", file.Filename))
		c.File("./" + file.Filename)
	})
	r.Run()
}
