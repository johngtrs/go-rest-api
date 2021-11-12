package utils

import (
	"io"
	"log"
	"os"
)

const (
	gopherSize   = 1934
	gloggerFile  = "glogger.log"
	tmpFile      = "tmp.log"
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func Glogger(prefix string, message string) {
	// Create temp file
	tmp, err := os.Create(tmpFile)
	if err != nil {
		panic(err)
	}

	defer tmp.Close()

	// Check if glogger file exists open and copy until the gopher
	// to the temp file
	if _, err := os.Stat(gloggerFile); err == nil {
		glogger, err := os.Open(gloggerFile)
		if err != nil {
			panic(err)
		}

		defer glogger.Close()

		totalBytes, err := glogger.Stat()
		if err != nil {
			panic(err)
		}

		io.CopyN(tmp, glogger, totalBytes.Size()-gopherSize)
	}

	// Add log in the file
	logger := log.New(tmp, prefix+" : ", log.LstdFlags)
	logger.Println(message)

	// Print log in the terminal
	log.Printf(ErrorColor, prefix+" : "+message)

	drawGoopher()
	replaceTmpFile()
}

func replaceTmpFile() {
	if _, err := os.Stat(gloggerFile); err == nil {
		if err := os.Remove(gloggerFile); err != nil {
			panic(err)
		}
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
