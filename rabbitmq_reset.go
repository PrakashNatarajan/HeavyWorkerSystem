package main

import(
	"log"
	"os/exec"
)

func main() {
	
	startrbtMQCmd := exec.Command("sudo", "service", "rabbitmq-server", "restart")
	log.Printf("Running restart command and waiting for it to finish...")
	stoprbtMQErr := startrbtMQCmd.Run()
	if stoprbtMQErr != nil {
		log.Printf("restart Command finished with error: %v", stoprbtMQErr)
	}
	stopappCmd := exec.Command("sudo", "rabbitmqctl", "stop_app")
	log.Printf("Running stop_app command and waiting for it to finish...")
	stopappErr := stopappCmd.Run()
	if stopappErr != nil {
		log.Printf("stop_app Command finished with error: %v", stopappErr)
	}
	resetCmd := exec.Command("sudo", "rabbitmqctl", "reset")
	log.Printf("Running reset command and waiting for it to finish...")
	resetErr := resetCmd.Run()
	if resetErr != nil {
		log.Printf("reset Command finished with error: %v", resetErr)
	}
	startappCmd := exec.Command("sudo", "rabbitmqctl", "start_app")
	log.Printf("Running start_app command and waiting for it to finish...")
	startappErr := startappCmd.Run()
	if startappErr != nil {
		log.Printf("start_app Command finished with error: %v", startappErr)
	}
	delvhostCmd := exec.Command("sudo", "rabbitmqctl", "delete_vhost", "/")
	log.Printf("Running delete_vhost command and waiting for it to finish...")
	delvhostErr := delvhostCmd.Run()
	if delvhostErr != nil {
		log.Printf("delete_vhost Command finished with error: %v", delvhostErr)
	}
	addvhostCmd := exec.Command("sudo", "rabbitmqctl", "add_vhost", "/")
	log.Printf("Running add_vhost command and waiting for it to finish...")
	addvhostErr := addvhostCmd.Run()
	if addvhostErr != nil {
		log.Printf("add_vhost Command finished with error: %v", addvhostErr)
	}
	guestpermsnsCmd := exec.Command("sudo", "rabbitmqctl", "set_permissions", "-p", "/", "guest", "^.*", ".*", ".*")
	log.Printf("Running set_permissions command and waiting for it to finish...")
	guestpermsnsErr := guestpermsnsCmd.Run()
	if guestpermsnsErr != nil {
		log.Printf("set_permissions Command finished with error: %v", guestpermsnsErr)
	}
	adduserCmd := exec.Command("sudo", "rabbitmqctl", "add_user", "admin", "admin")
	log.Printf("Running add_user command and waiting for it to finish...")
	adduserErr := adduserCmd.Run()
	if adduserErr != nil {
		log.Printf("add_user Command finished with error: %v", adduserErr)
	}
	admintagsCmd := exec.Command("sudo", "rabbitmqctl", "set_user_tags", "admin", "administrator")
	log.Printf("Running set_user_tags command and waiting for it to finish...")
	adminusrtagsErr := admintagsCmd.Run()
	if adminusrtagsErr != nil {
		log.Printf("set_user_tags Command finished with error: %v", adminusrtagsErr)
	}
	adminpermsnsCmd := exec.Command("sudo", "rabbitmqctl", "set_permissions", "-p", "/", "admin", ".*", ".*", ".*")
	log.Printf("Running set_permissions command and waiting for it to finish...")
	adminpermsnsErr := adminpermsnsCmd.Run()
	if adminpermsnsErr != nil {
		log.Printf("set_permissions Command finished with error: %v", adminpermsnsErr)
	}

}