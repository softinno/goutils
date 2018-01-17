package sockt

import (
	"fmt"
	"io"
	"io/ioutil"
    "net"
    "os"
    "strings"
    "strconv"
)

// Create a socket
// - param: service = host:port
// Return: tcp connection
func CreateConn(service string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    if check(err) {
	    return nil, err
    }
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if check(err) {
	    return conn, err
    }
    return conn, nil
}

// Do tunnel 
// 
func DoTunnel(conn *net.TCPConn, host string, port int, user string, pass string) (*net.TCPConn, error) {
	portStr := strconv.Itoa(port)
	msg := "CONNECT " + host + ":" + portStr + 
	    " HTTP/1.0\r\nHost: " + host + ":" + portStr + 
        "\r\nUser-Agent: Mozilla/4.0 (compatible; MSIE 6.0; Win32)\r\n"
        //+ "Proxy-Connection: Keep-Alive\r\n";
        if (isAuthValid(user, pass)) {
        	msg += "Proxy-Authorization: Basic "+ user +":" + pass +"\r\n"
        }
        msg += "\r\n"
        
	fmt.Println("SEND \n" + msg)
	_, err := conn.Write([]byte(msg))
	fmt.Println("SEND DONE \n")
    if check(err) {
    	return conn, err
    }
    
    
    fmt.Println("RECV \n")
		
	reply := make([]byte, 600)
	tmp := make([]byte, 1)
	replyLen := 0
	newlinesSeen := 0
	headerDone := false

	// read all up to 2 consecutive \n, save only header (first line) 
    for (newlinesSeen < 2) {
	    _, err := conn.Read(tmp)
	    if err != nil {
            if err != io.EOF {
                fmt.Println("read error:", err)
            }
            break
        }	
        //fmt.Println("RECV n=" + strconv.Itoa(n) + " c=" + string(tmp));
	    if (tmp[0] == '\n') {
            headerDone = true;
            newlinesSeen++;
        } else if (tmp[0] != '\r') {
            newlinesSeen = 0;
            if (!headerDone && replyLen < len(reply)) {
            	reply[replyLen] = tmp[0];
            	replyLen++
            }
        }
    }
    
    resultStr := string(reply)
    fmt.Println("RECV DONE \n"+ resultStr)
    
    if (! (strings.HasPrefix(resultStr, "HTTP/1.1 200") || 
    	   strings.HasPrefix(resultStr, "HTTP/1.0 200"))) {
	    fmt.Println("Unable to tunnel to " + host + portStr)
    }
    
    return conn, err
}

// Check http server
// - param: service = host:port
func CheckHttpSrv(service string) error {
	conn, err := CreateConn(service) 
	if 
	check(err) {
    	return err
    }
    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    if check(err) {
    	return err
    }
    
    result, err := ioutil.ReadAll(conn)
    if check(err) {
    	return err
    }
    
    fmt.Println(string(result))
    return nil
}

func Get(conn *net.TCPConn, path string) (string, error) {
    _, err := conn.Write([]byte("GET " + path + " HTTP/1.0\r\n\r\n"))
    if check(err) {
    	return "", err
    }
    
    result, err := ioutil.ReadAll(conn)
    if check(err) {
    	return string(result), err
    }
    
    //fmt.Println(string(result))
    return string(result), nil
}

// ---
func isAuthValid(user string, pass string) bool {
	if (user != "") {
		return true
	} else {
		return false
	}
}

func check(err error) bool {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
        return true
    }
    return false
}
