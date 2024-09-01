package main                                                                  
                                                                                  
    import (                                                                      
    	"fmt"                                                                         
    	"log"                                                                         
    	"os"                                                                          
    	"os/signal"                                                                   
    	"path/filepath"                                                               
    	"syscall"                                                                     
    	"time"                                                                        
    	"context"                                                                     
    	"github.com/johnconnor-sec/gsv-go/main" // Adjust this import path based on actual structure            
    )                                                                             
                                                                                  
    const lockFile = "/tmp/run_script.lock"                                       
                                                                                  
    func createLockFile() error {                                                 
    	if _, err := os.Stat(lockFile); err == nil {                                  
    		return fmt.Errorf("script is already running")                                
    	}                                                                             
    	file, err := os.Create(lockFile)                                              
    	if err != nil {                                                               
    		return err                                                                    
    	}                                                                             
    	file.Close()                                                                  
    	return nil                                                                    
    }                                                                             
                                                                                  
    func removeLockFile() {                                                       
    	if err := os.Remove(lockFile); err != nil {                                   
    		log.Printf("Failed to remove lock file: %v", err)                             
    	}                                                                             
    }                                                                             
                                                                                  
    func setupSignalHandler() {                                                   
    	sigs := make(chan os.Signal, 1)                                               
    	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)                          
    	go func() {                                                                   
    		<-sigs                                                                        
    		removeLockFile()                                                              
    		os.Exit(1)                                                                    
    	}()                                                                           
    }                                                                             
                                                                                  
    func main() {                                                                 
    	// Setup lock file and signal handler                                         
    	if err := createLockFile(); err != nil {                                      
    		log.Fatalf("%v", err)                                                         
    	}                                                                             
    	defer removeLockFile()                                                        
    	setupSignalHandler()                                                          
                                                                                  
    	// Setup logging                                                              
    	logDir := filepath.Join(os.Getenv("HOME"),                                    
  "tools/my_youtube_playlists/get_saved_videos/gsv-go/logs")                             
    	if err := os.MkdirAll(logDir, 0755); err != nil {                             
    		log.Fatalf("Failed to create log directory: %v", err)                         
    	}                                                                             
    	logFile := filepath.Join(logDir, "script.log")                                
    	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)     
    	if err != nil {                                                               
    		log.Fatalf("Failed to open log file: %v", err)                                
    	}                                                                             
    	defer f.Close()                                                               
    	log.SetOutput(f)                                                              
                                                                                  
    	log.Printf("%s - Starting script", time.Now().Format("2006-01-02 15:04:05"))  
                                                                                  
    	// Run main logic from gsv-go/main.go                                         
    	ctx := context.Background()                                                   
    	if err := main.Run(ctx); err != nil {                                         
    		log.Fatalf("Error running gsv-go: %v", err)                                   
    	}                                                                             
                                                                                  
    	log.Printf("%s - Script finished", time.Now().Format("2006-01-02 15:04:05"))  
                                                                                  
    	// Send notification                                                          
    	if err := exec.Command("notify-send", "YouTube Playlist Update", "The         
  playlist update script has finished running.").Run(); err != nil {              
    		log.Printf("Failed to send notification: %v", err)                            
    	}                                                                             
    } 
