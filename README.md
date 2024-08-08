# Amparo

## Getting Started

### Prerequisites

- **Go**: Make sure Go is installed on your machine. You can download it from [here](https://golang.org/dl/).
- **Make**: Ensure you have `make` installed, as the project uses a `Makefile` for easy execution of commands.
- **AWS CLI**: If deploying to AWS Lambda, you should have the AWS CLI installed and configured.

### Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/caiquetorres/amparo.git
    cd amparo
    ```

2. **Create a `.env` file**:
    - Duplicate the `.env.example` file and rename it to `.env`.
    - Modify the `.env` file as needed for your environment.

```bash
cp .env.example .env
```

3. **Run the app locally**:
    - Use the following command to run the application locally:

```bash
make run
```

### Testing

To run tests for the application:

```bash
make test
```
