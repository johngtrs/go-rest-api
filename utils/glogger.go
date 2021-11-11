package utils

import (
	"io"
	"log"
	"os"
)

const gopherSize = 1934
const gloggerFile = "glogger.log"
const tmpFile = "tmp.log"

func Glogger(prefix string, message string) {
	glogger, err := os.Open(gloggerFile)
	if err != nil {
		panic(err)
	}
	defer glogger.Close()

	dest, err := os.Create(tmpFile)
	if err != nil {
		panic(err)
	}

	defer dest.Close()

	totalBytes, err := glogger.Stat()
	if err != nil {
		panic(err)
	}

	io.CopyN(dest, glogger, totalBytes.Size()-gopherSize)
	logger := log.New(dest, prefix+" : ", log.LstdFlags)
	logger.Println(message)

	drawGoopher()
	replaceTmpFile()
}

func replaceTmpFile() {
	if err := os.Remove(gloggerFile); err != nil {
		panic(err)
	}

	if err := os.Rename(tmpFile, gloggerFile); err != nil {
		panic(err)
	}
}

func drawGoopher() {
	file, err := os.OpenFile(tmpFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	file.WriteString(`

                         .:-==+++++++++++++++++++++=-:                  
                    .-+**+=-:.                      .:-+**=.                    
                 -**+:                           .:::.    .=*+:.-+++++++-       
    -+++++++++-**:   .-::::::-:              :---.   .:--:   .+%+       -**     
  -#-       :#=   .--.        .--:         -=.            --    =#. .     -#    
 -%     .. +#.   --              :=       +.               .+    .%@@@=    #-   
 #-    +@@@+    +.                 +     =. :==-.            =     #@@-    *=   
 *+    -@@=    .+.*@@@%-           :-    + #@@@@@-           =      #+    -%    
 .%-    *+     --#@@@@*@           .+    +.@@@@=-+           =       %: -*+     
   +#=.-%      .*-%@@%==           -:    :-.*%%*-           .=       -@*-       
     :=@-       :- .::            .=      :-               :=         #=        
      .%         :=.             --.=*###*-.--.          :-:          .%        
      +*           :--:      .--- :@@@@@@@@+  :----------              #-       
      %-               ::::::.  --:-*#%%%*-.--:                        +*       
     .@:                       =              .=                       -%       
     .@.                       =     :---.     =                       :@       
     :@                         :::*:  + .-+::-                        .@.      
     -@                            =   +   -.                           @.      
     :@                           .=   #   -.                           @:      
     .@.                           -::-:=::-                            @:      
      @:                                                                @:      
      #=                                                                @:      
      =*                                                                @:      
       %                                                                @:   `)

}
