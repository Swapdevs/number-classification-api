# Number Classification API

An API that takes a number and returns interesting mathematical properties about it, along with a fun fact. Built with Go, containerized using Docker, and deployed on AWS Elastic Beanstalk.

## Features
The API classifies a number and returns:
- Whether it is **prime**.
- Whether it is **perfect**.
- Properties like **"armstrong"**, **"odd"**, or **"even"**.
- **Sum of its digits**.
- A **fun fact** about the number.
- Handles **CORS (Cross-Origin Resource Sharing)**.
- Returns responses in **JSON format**.

## Technologies Used
- **Programming Language:** Go (Golang)
- **Containerization:** Docker
- **Deployment:** AWS Elastic Beanstalk
- **Version Control:** GitHub

## API Endpoint
```
GET /api/classify-number?number=<number>
```

### Example Request
```bash
curl "http://localhost:8080/api/classify-number?number=371"
```

### Example Response
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```

## Running the Project Locally
### Prerequisites
- Go (1.19 or higher)
- Docker (optional, for containerization)

### Steps
1. **Clone the repository:**
   ```bash
   git clone https://github.com/Swapdevs/number-classification-api.git
   ```
2. **Navigate to the project directory:**
   ```bash
   cd number-classification-api
   ```
3. **Run the application:**
   ```bash
   go run main.go
   ```
4. **Access the API at:**
   ```
   http://localhost:8080/api/classify-number?number=371
   ```

## Running with Docker
1. **Build the Docker image:**
   ```bash
   docker build -t number-classification-api .
   ```
2. **Run the Docker container:**
   ```bash
   docker run -p 8080:8080 number-classification-api
   ```
3. **Access the API at:**
   ```
   http://localhost:8080/api/classify-number?number=371
   ```
![](https://github.com/Swapdevs/number-classification-api/blob/main/images/lt.png)
## Deployment
The API is deployed using **AWS Elastic Beanstalk**. You can access it at:
```
http://number-classification-api-env.eba-cmnsufps.us-east-1.elasticbeanstalk.com/api/classify-number?number=371
```

### Steps to Deploy to AWS
1. **Push the Docker image to Amazon ECR.**
2. **Create an Elastic Beanstalk environment.**
3. **Configure the environment to use the Docker image.**
4. **Access the API using the provided Elastic Beanstalk URL.**

## Project Structure
```
number-classification-api/
├── main.go               # Main Go application code
├── Dockerfile            # Docker configuration
├── README.md             # Project documentation
├── go.mod                # Go module file
└── go.sum                # Go dependencies checksum
```

## API Response Format
### Success Response (200 OK)
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```

### Error Response (400 Bad Request)
```json
{
    "number": "alphabet",
    "error": true
}
```
![](https://github.com/Swapdevs/number-classification-api/blob/main/images/dpt.png)
![](https://github.com/Swapdevs/number-classification-api/blob/main/images/webt.png)
## Contributing
Contributions are welcome! If you'd like to contribute, please follow these steps:

1. **Fork the repository.**
2. **Create a new branch:**
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. **Commit your changes:**
   ```bash
   git commit -m "Add your feature"
   ```
4. **Push to the branch:**
   ```bash
   git push origin feature/your-feature-name
   ```
5. **Open a pull request.**

## License
This project is licensed under the **MIT License**. Please take a look at the [LICENSE](LICENSE) file for details.

## Acknowledgments
- Inspired by the need to learn and explore **Go, Docker, and AWS**.
- Special thanks to the **Go community** for their excellent documentation and resources.

## Contact
If you have any questions or suggestions, feel free to reach out:
- **GitHub:** [Swapdevs](https://github.com/swapdevs)
- **Email:** christopher.akinsanmi@gmail.com

Enjoy exploring the **Number Classification API**! 🚀

