package clients

type Client struct {
	ID       int
	Name     string
	Password string
	Balance  int
}

// CREATE TABLE `clients` (
//     `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
//     `login_id` int NOT NULL,
//     -- `password` varchar(300) NOT NULL,
//     `balance` int NOT NULL
// );
