package controller

import (
	"SQLInjection/dto"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

type ExecutionController struct{

}


func (ec *ExecutionController) PingTheHost (c *gin.Context){
	sess := sessions.Default(c)
	var pRequest dto.PingRequest
	if err := c.ShouldBind(&pRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request !",
		})
		return
	}
	if sess.Get("id") != nil {
		argument := fmt.Sprintf("ping -c 1 %s",pRequest.Host)
		out,err := exec.Command("/bin/bash", "-c",argument).Output()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Command Execution Failed !",
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Command Execution Successful !",
			"output": out,
		})
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Please Login",
		})
	}
}

func (ec *ExecutionController) PingTheHost2 (c *gin.Context){
	sess := sessions.Default(c)
	var pRequest dto.PingRequest
	if err := c.ShouldBind(&pRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request !",
		})
		return
	}
	if sess.Get("id") != nil {
		pRequest.Host = strings.Replace(pRequest.Host, ";","",1)
		argument := fmt.Sprintf("ping -c 1 %s",pRequest.Host)
		out,err := exec.Command("/bin/bash", "-c",argument).Output()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Command Execution Failed !",
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Command Execution Successful !",
			"output": out,
		})
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Please Login",
		})
	}
}

func (ec *ExecutionController) PingTheHost3 (c *gin.Context){
	sess := sessions.Default(c)
	var pRequest dto.PingRequest
	if err := c.ShouldBind(&pRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request !",
		})
		return
	}
	if sess.Get("id") != nil {
		pRequest.Host = strings.Replace(pRequest.Host, ";","",-1)
		argument := fmt.Sprintf("ping -c 1 %s",pRequest.Host)
		out,err := exec.Command("/bin/bash", "-c",argument).Output()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Command Execution Failed !",
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Command Execution Successful !",
			"output": out,
		})
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Please Login",
		})
	}
}

func (ec *ExecutionController) PingTheHost4 (c *gin.Context){
	sess := sessions.Default(c)
	var pRequest dto.PingRequest
	if err := c.ShouldBind(&pRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request !",
		})
		return
	}
	if sess.Get("id") != nil {
		blackChars := []string{";", "`"}
		for _, char := range blackChars{
			pRequest.Host = strings.Replace(pRequest.Host, char, "",-1)
		}
		argument := fmt.Sprintf("ping -c 1 %s",pRequest.Host)
		out,err := exec.Command("/bin/bash", "-c",argument).Output()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Command Execution Failed !",
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Command Execution Successful !",
			"output": out,
		})
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Please Login",
		})
	}
}

func (ec *ExecutionController) PingTheHost5 (c *gin.Context){
	sess := sessions.Default(c)
	var pRequest dto.PingRequest
	if err := c.ShouldBind(&pRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request !",
		})
		return
	}
	if sess.Get("id") != nil {
		blackChars := []string{";", "`", "$", "(", ")"}
		for _, char := range blackChars{
			pRequest.Host = strings.Replace(pRequest.Host, char, "",-1)
		}
		argument := fmt.Sprintf("ping -c 1 %s",pRequest.Host)
		out,err := exec.Command("/bin/bash", "-c",argument).Output()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Command Execution Failed !",
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Command Execution Successful !",
			"output": out,
		})
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Please Login",
		})
	}
}


func (ec *ExecutionController) ExecuteCommandFile (c *gin.Context){
	sess := sessions.Default(c)
	if sess.Get("id") != nil {
		filename := "/var/temp/" + c.Param("id") + ".txt"
		content, err := ioutil.ReadFile(filename)
		if err != nil{
			c.JSON(http.StatusNotFound, gin.H{
				"message": "File Not Exist !",
			})
			return
		}
		command := string(content)
		command = strings.Replace(command, "\n", "", -1)
		command = strings.Replace(command, "\t","", -1)
		out,err := exec.Command(command).Output()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Command Execution Failed !",
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Command Execution Successful !",
			"output": out,
		})
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Please Login",
		})
	}
}