package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"../gethTest"
)

var (
	targetIp   string
	targetPort int
	ourUdpPort int
	ourTcpPort int

	privKeyFile   string
	targetKeyFile string

	keyDir string
)

func init() {
	// flag.StringVar(&targetIp, "tip", "116.202.237.156", "target IP address")
	flag.StringVar(&targetIp, "tip", "172.18.1.44", "target IP address")
	flag.IntVar(&targetPort, "tport", 30303, "target port number")
	// flag.IntVar(&ourUdpPort, "oudpport", 30312, "our udp port number")
	// flag.IntVar(&ourTcpPort, "otcpport", 30303, "our tcp port number")
	flag.IntVar(&ourUdpPort, "oudpport", 30400, "our udp port number")
	flag.IntVar(&ourTcpPort, "otcpport", 30401, "our tcp port number")
	flag.StringVar(&privKeyFile, "keyfile",
		"../keys/0_0", "Private key file")
	flag.StringVar(&targetKeyFile, "targetkeyfile",
		"../key", "Target public key file")
	flag.StringVar(&keyDir, "keydir",
		"../keys", "Key directory")
}

// extractTargetNodeId fetches the NodeID of the specified target
// and writes it to disk
func extractTargetNodeId() {
	s := gethTest.NewPingServer(targetIp, targetPort, ourUdpPort, ourTcpPort)
	s.GeneratePrivateKey()
	s.Start()
	time.Sleep(2 * time.Second)
	s.Stop()
	s.WriteTargetIdFile(targetKeyFile)
}

// generateKeys generates the specified number of keys
// for every bucket of the target
func generateKeys(num int) {
	s := gethTest.NewPingServer(targetIp, targetPort, ourUdpPort, ourTcpPort)
	total := 17 * num // as of 09.2018, the number of buckets in geth is set to 17

	k := gethTest.NewKeyStore(num)
	s.GeneratePrivateKey()
	s.ParseTargetIdFile(targetKeyFile)
	for k.KeysTotal() < total {
		s.GeneratePrivateKey()
		foo := s.PrivateKey()
		bucketNum := s.BucketNumber()
		k.Add(foo, bucketNum)
	}
	k.WriteKeysToFolder(keyDir)
}

// pingLoopRand starts a ping loop using a randomly
// generated private key and NodeID
func pingLoopRand() {
	s := gethTest.NewPingServer(targetIp, targetPort, ourUdpPort, ourTcpPort)
	s.GeneratePrivateKey()
	s.Start()
	time.Sleep(72 * time.Hour)
	s.Stop()
}

// pingLoop starts a ping loop using the private key
// specified in the private key file
func pingLoop() {
	s := gethTest.NewPingServer(targetIp, targetPort, ourUdpPort, ourTcpPort)
	s.ParsePrivateKeyFile(privKeyFile)
	// s.GeneratePrivateKey()
	s.Start()
	time.Sleep(1 * time.Second)
	s.Stop()
}

func pingLoopFromFile(filename string) {
	s := gethTest.NewPingServer(targetIp, targetPort, ourUdpPort, ourTcpPort)
	s.ParsePrivateKeyFile(filename)
	// s.GeneratePrivateKey()
	s.Start()
	time.Sleep(1 * time.Second)
	s.Stop()
}

func launchAttack() {
	for i := 0; i < 6; i++ {
		for j := 2; j < 4; j++ {
			pingLoopFromFile("../keys/" + strconv.Itoa(j) + "_" + strconv.Itoa(i))
			fmt.Println("../keys/" + strconv.Itoa(j) + "_" + strconv.Itoa(i))
			time.Sleep(1 * time.Second)
		}
	}
}

func getLogdistFromFile(filename string) {
	r, _ := os.Open(filename)
	distList := []int{}
	defer r.Close()
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		victimString := "a1ba4012cb2d0477e144dad0071e4175f2db8b9ccdd85e8a3ee8be20a067acb6a250773c6b9f5e8b92ad544098481284715f5410159ffa83362440f79580df53"
		victim, _ := hex.DecodeString(victimString)
		attack, _ := hex.DecodeString(line)
		distList = append(distList, gethTest.Bucket(victim, attack))
	}
	fmt.Println(distList)
}

func main() {
	// launchAttack()
	getLogdistFromFile(("/home/ensdaddy/packetsdata/norm_pubkey"))
	// extractTargetNodeId()
	// generateKeys(25)
	//---------------------caculate logdist of TargetNode and victim---------------------
	// victimString := "84cbb6039ac0a5a16ce3bd41d075906ba7bf19268f1a1b4d3568275145f15c1e1c0aaaa9eb505e38fe231603c6c63a074e81b942fdae53a8bd5a12f4f639a082"
	// victim, _ := hex.DecodeString(victimString)
	// strings := []string{
	// 	"5c0342870fba86b8f33523f604b2abca06a1df45b93f747b4aee5a93061c79b3331ebda4fc9df781605c8c5e24caeb638ec0171aff9f6e50217277c11743467d",
	// }
	// for _, attack := range strings {
	// 	attackId, _ := hex.DecodeString(attack)
	// 	fmt.Println(gethTest.Bucket(attackId, victim))
	// }
	// ----------------------------------------------------------------
}
