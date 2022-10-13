package main

import "bufio"
import "os"
import "fmt"
import "strings"
import "os/exec"

func main() {
  
  reader := bufio.NewReader(os.Stdin)
  
  for {
    
    fmt.Print("> ")
    input, err := reader.ReadString('\n')
    
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
    }

    if err = execInput(input); err != nil {
      fmt.Fprintln(os.Stderr, err)
    }
  }
}

func execInput(input string) error {
  input = strings.TrimSuffix(input, "\n")
  args  := strings.Split(input, " ")  
  
  switch args[0] {
  case "cd":
    if len(args) < 2 {
      home_dir, _ := os.UserHomeDir()
      return os.Chdir(home_dir)
    }
    return os.Chdir(args[1])
  case "pwd":
    current_dir, _ := os.Getwd()
    fmt.Println(current_dir)    
    return nil
  case "hostname":
    current_user, _ := os.Hostname()
    fmt.Println(current_user)
    return nil
  case "exit":
    os.Exit(0)
  }

  cmd := exec.Command(args[0], args[1:]...)
  
  cmd.Stderr = os.Stderr
  cmd.Stdout = os.Stdout
  
  return cmd.Run()
} 
