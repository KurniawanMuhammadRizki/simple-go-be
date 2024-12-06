# Simple Go Backend

It is written in Go, using GORM for ORM, Fiber for the web framework, and PostgreSQL as the database. Docker is used to run both the PostgreSQL instance and the Go application.

## Prerequisites

- Go 1.21.3 or later
- Docker
- Docker Compose

## Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/KurniawanMuhammadRizki/simple-go-be.git
   ```

2. Create .env file:
  At the root of the project, create a .env file with the following content:

   <img width="148" alt="image" src="https://github.com/user-attachments/assets/0723e765-fad4-4a64-bb5b-0a92bfa94945">


### .env 

    ```env
    # Database connection
    POSTGRES_DB=postgres
    POSTGRES_HOST=postgres
    POSTGRES_PASSWORD=admin
    POSTGRES_PORT=5432
    POSTGRES_USER=admin
    
    # Database pool configuration
    POSTGRES_IDLE_CONNECTION=10
    POSTGRES_MAX_CONNECTION=100
    POSTGRES_MAX_LIFETIME_CONNECTION=300 # in seconds
    
    # Application
    APP_PORT=8080
    APP_NAME=simple-go-be
    
    # Logging
    LOG_LEVEL=4
    ```

3. Start Docker:
  Run the following command to start Docker. This will initialize database migrations, spin up a PostgreSQL instance, and start the Go application
   ```sh
   docker-compose up -d
   ```
  If everything runs successfully, you should see an output similar to this:
  <img width="764" alt="image" src="https://github.com/user-attachments/assets/373d2d74-2bed-44ac-b775-6a5f03f0e099">


## Project Structure

- `go.mod` - Go module file
- `go.sum` - Go dependencies checksum file
- `compose.yml` - Docker Compose file for PostgreSQL
- `.env` - Environment variables file
- `.gitignore` - Git ignore file

## Dependencies

- [GORM](https://gorm.io/)
- [Fiber](https://gofiber.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)

## API Examples

Here are some examples of how to use the API:

### 1. Create a Brand

**Endpoint:**  
`POST http://127.0.0.1:8080/brands/`  

**Request Body:**
```json
{
  "name": "Samsung"
}
```
### 2. Create a Voucher

**Endpoint:**  
`POST http://127.0.0.1:8080/vouchers/`  

**Request Body:**
```json
{
    "name": "Voucher Discount Rp. 10.000",
    "brand_id": 1,
    "cost_in_point": 10000
}
```
### 3. Get Voucher By ID

**Endpoint:**  
`GET http://127.0.0.1:8080/vouchers/1`  

### 4. Get Voucher By Brand

**Endpoint:**  
`GET http://127.0.0.1:8080/vouchers/brand?id=1` 

### 5. Create a Customer

**Endpoint:**  
`POST http://127.0.0.1:8080/customers/`  

**Request Body:**
```json
{
     "name": "Fulan"
}
```

### 6. Get Customer By ID

**Endpoint:**  
`GET http://127.0.0.1:8080/customers/1` 

### 7. Create a Redemption

**Endpoint:**  
`POST http://127.0.0.1:8080/transaction/redemption`  

**Request Body:**
```json
{
    "customer_id": 1,
    "voucher_items": [
        {
            "voucher_id": 1,
            "quantity": 5
        },
        {
            "voucher_id": 3,
            "quantity": 5
        }
    ]
}
```
### 8. Get Transaction Detail

**Endpoint:**  
`GET http://127.0.0.1:8080/transaction/redemption/transactionID?id=8` 

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
