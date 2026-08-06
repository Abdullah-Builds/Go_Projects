# 📧 Email Dispatcher

A concurrent email dispatching service built with **Go** that sends emails efficiently using worker pools, goroutines, and channels. The project is designed to handle bulk email delivery while maintaining high performance and clean architecture.

## ✨ Features

- Concurrent email sending using goroutines
- Worker pool implementation for controlled concurrency
- Job queue with Go channels
- SMTP email support
- HTML/Text email templates
- CSV or structured recipient input
- Error handling and logging
- Configurable number of workers
- Clean and modular project structure

## 🛠️ Tech Stack

- Go (Golang)
- SMTP
- Goroutines
- Channels
- Worker Pools
- HTML Templates

## 📂 Project Structure

```
Email-Dispatcher/
│── cmd/
│── config/
│── internal/
│   ├── dispatcher/
│   ├── email/
│   ├── worker/
│   └── templates/
│── recipients/
│── logs/
│── go.mod
│── go.sum
└── main.go
```

## ⚙️ How It Works

1. Load recipient data.
2. Create email jobs.
3. Push jobs into a buffered channel.
4. Spawn multiple worker goroutines.
5. Each worker reads jobs from the queue.
6. Workers send emails through the configured SMTP server.
7. Successes and failures are logged.

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- SMTP account (Gmail, Mailtrap, SendGrid, etc.)

### Clone the Repository

```bash
git clone https://github.com/Abdullah-Builds/Go_Projects.git
cd Go_Projects
```

### Install Dependencies

```bash
go mod tidy
```

### Configure SMTP

Update your SMTP credentials in the configuration file or environment variables.

Example:

```env
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_EMAIL=your-email@gmail.com
SMTP_PASSWORD=your-app-password
```

### Run the Application

```bash
go run main.go
```

## Example Workflow

```
Recipients
      │
      ▼
 Job Queue (Channel)
      │
      ▼
 Worker Pool
 ┌────┼────┐
 │    │    │
 ▼    ▼    ▼
SMTP SMTP SMTP
 │    │    │
 ▼    ▼    ▼
Emails Sent
```

## Configuration

You can customize:

- Number of worker goroutines
- SMTP server
- Email templates
- Retry logic
- Batch size
- Timeout settings

## Future Improvements

- Retry mechanism with exponential backoff
- Email scheduling
- Attachment support
- Rate limiting
- REST API
- Docker support
- Metrics and monitoring
- Email tracking

## Learning Objectives

This project demonstrates:

- Goroutines
- Channels
- Worker Pools
- Concurrency patterns
- SMTP integration
- Go project organization
- Error handling
- Logging

## Contributing

Contributions are welcome.

1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Open a Pull Request.

## License

This project is licensed under the MIT License.

## Author

**Abdullah Builds**

GitHub: https://github.com/Abdullah-Builds
