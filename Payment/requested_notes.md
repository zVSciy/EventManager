# Pipeline Documentation

## GitHub Actions CI Pipeline

The Payment Service uses GitHub Actions for continuous integration. The pipeline is defined in `.github/workflows/go.yml`.

### Pipeline Configuration

```yaml
name: Go

on:
  push:
    paths:
      - 'Payment/**'
  pull_request:
    paths:
      - 'Payment/**'
```

This section configures the pipeline triggers:
- Runs on any push that modifies files in the `Payment/` directory
- Runs on any pull request that modifies files in the `Payment/` directory
- Changes to other directories won't trigger this pipeline

### Environment and Setup

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.23'
        cache: false
```

This section:
- Uses Ubuntu latest as the runner environment
- Checks out the repository using actions/checkout@v4
- Sets up Go version 1.23 using actions/setup-go@v4

### Caching Configuration

```yaml
    - name: Cache Go modules
      uses: actions/cache@v4
      with:
        path: |
          ~/.cache/go-build
          ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('Payment/src/api/go.sum') }}
        restore-keys: |
          ${{ runner.os }}-go-
```

The caching strategy:
- Caches Go build cache and module cache
- Cache key is based on:
  - Operating system
  - Contents of `Payment/src/api/go.sum`
- Falls back to previous caches if exact match isn't found
- Significantly speeds up subsequent builds

### Build and Test Steps

```yaml
    - name: Build
      working-directory: Payment/src/api
      run: go build -v ./...

    - name: Test
      working-directory: Payment/src/api
      run: go test -v ./...
```

These steps:
1. Build Process
   - Runs in `Payment/src/api` directory
   - Builds all packages recursively
   - Uses verbose output (`-v` flag)

2. Test Execution
   - Runs in the same directory
   - Executes all tests recursively
   - Uses verbose output for detailed test results

### Pipeline Status Requirements

The pipeline must complete successfully for:
- All push events to be accepted
- Pull requests to be eligible for merging

<br>

# What's new?

## Authentication System

The Payment Service currently implements a temporary authentication mechanism pending integration with the User Management service.

### Current Implementation Details

The authentication is implemented in `internal/util/auth_utils.go` with the following characteristics:

#### Token Validation
- Every request (except `/health`) must include an `Authorization` header
- The header must be in the format: `Bearer <token>`
- **Important Note**: Currently, the service accepts ANY non-empty Bearer token as valid
- This is a temporary implementation until integration with the User Management service is complete

#### Endpoint Security
All secured endpoints validate:
1. Presence of the Authorization header
2. Token format (must be Bearer)
3. User ID matching (token must correspond to the user making the request)

```go
// Example from auth_utils.go
func ValidateAuthHeader(userID, authHeader string) error {
    if authHeader == "" {
        return errors.New("missing auth header")
    }
    // Currently returns nil for any non-empty auth header
    return nil
}
```

### Security Limitations

Current known limitations of the authentication system:
1. No token signature validation
2. No token expiration checking
3. No verification against User Management service
4. Any well-formed Bearer token is accepted
5. No role-based access control

### Usage in API Endpoints

Every secured endpoint requires:
1. Authorization header with Bearer token
2. Token must match the user ID in the request path or body

Example endpoints and their auth requirements:
- `GET /accounts/{user_id}` - Token must match user_id
- `POST /payments` - Token must match payment.user_id
- `GET /accounts/{user_id}/payments` - Token must match user_id

### Error Responses

When authentication fails, endpoints return:
```json
{
    "error": "Unauthorized"
}
```
with HTTP status code 401.

## New & modified test cases

Authentication introduced new behavior and edge cases, prompting the addition of new tests and revisions to some existing ones.