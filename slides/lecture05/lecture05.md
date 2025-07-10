---
header: "Summer 2025 Golang + Flutter Course: Rapid cross-platform development"
footer: "Lecture 05: Advanced Patterns"
title: "Lecture 05: Advanced Patterns"
author: "Timur Harin"
date: "2025-07-16"
theme: uncover
paginate: true
style: |
  section {
    background: #ffffff;
    color: #1a1a1a;
    font-size: 24px;
    line-height: 1.2;
    padding: 40px;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    height: 100vh;
    box-sizing: border-box;
  }
  section h1 {
    font-size: 60px;
    margin-bottom: 10px;
    margin-top: 60px;
    font-weight: 700;
    text-align: center;
    flex-shrink: 0;
  }
  section h2 {
    font-size: 50px;
    margin-bottom: 14px;
    font-weight: 600;
  }
  section h3 {
    font-size: 40px;
    margin-bottom: 10px;
    font-weight: 600;
  }
  section h4 {
    font-size: 30px;
    margin-bottom: 8px;
    font-weight: 300;
  }
  section code {
    background: #f8f9fa;
    padding: 3px 6px;
    border-radius: 4px;
    color: #2563eb;
    font-family: 'JetBrains Mono', 'SF Mono', 'Monaco', monospace;
    font-size: 0.9em;
  }
  section pre {
    background: #f8f9fa;
    border: 2px solid #e5e7eb;
    border-radius: 12px;
    padding: 20px;
    margin: 15px 0;
    overflow-x: auto;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
    position: relative;
  }
  section pre code {
    background: none;
    padding: 0;
    color: #1f2937;
    font-family: 'JetBrains Mono', 'SF Mono', 'Monaco', monospace;
    font-size: 16px;
    line-height: 1.5;
    display: block;
  }
  .code-comment {
    color: #6b7280;
    font-style: italic;
  }
  section blockquote {
    background: #f8fafc;
    border-left: 4px solid #3b82f6;
    padding: 16px 24px;
    margin: 16px 0;
    border-radius: 6px;
    font-style: italic;
  }
  section ul, section ol {
    text-align: left;
    margin-left: 40px;
    color: #374151;
    margin-top: 5px;
    margin-bottom: 5px;
  }
  section li {
    margin-bottom: 3px;
  }
  section strong {
    color: #1f2937;
    font-weight: 700;
  }
  section table {
    border-collapse: collapse;
    margin: 20px 0;
    width: 100%;
    background: white;
    border-radius: 8px;
    overflow: hidden;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  }
  section th, section td {
    border: 1px solid #e5e7eb;
    padding: 12px 16px;
    text-align: left;
  }
  section th {
    background: #f9fafb;
    color: #374151;
    font-weight: 600;
  }
  .code-columns {
    display: flex;
    gap: 20px;
    align-items: flex-start;
  }
  .code-columns > div {
    flex: 1;
  }
  .code-columns pre {
    margin: 0;
  }
  .slide-content {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    text-align: left;
  }
  .slide-title {
    text-align: center;
    flex-shrink: 0;
  }
  section > h2:first-of-type,
  section > h3:first-of-type,
  section > p:first-of-type,
  section > ul:first-of-type,
  section > ol:first-of-type,
  section > pre:first-of-type,
  section > blockquote:first-of-type,
  section > div:first-of-type {
    margin-top: 0;
  }
---

<div class="slide-title">

# Go + Flutter Course
## Advanced Patterns

**Timur Harin**  
Lecture 05: **Advanced Patterns**

</br>

*Building scalable, maintainable, and secure applications*

</div>

---

# Block 5: Advanced Patterns

<div class="slide-content">

## Lecture 05 Overview
- **Clean Architecture**: Domain-driven design principles
- **Testing Strategies**: Unit, integration, and end-to-end testing
- **JWT & Authentication**: Secure user management
- **Security Best Practices**: Protecting applications from threats

## What we'll learn
- **Why clean architecture matters for long-term maintenance**
- **Dependency inversion and separation of concerns**
- **Comprehensive testing strategies for both Go and Flutter**
- **JWT implementation with proper security measures**
- **Common security vulnerabilities and prevention**
- **Production-ready security practices**
- **Performance and monitoring considerations**

</div>

---

# Learning path

<div class="slide-content">

```
Advanced Development Patterns
       |
       ├── Clean Architecture ─────────┐
       |                               |
       ├── Testing Strategies ─────────┤
       |                               |
       ├── JWT & Authentication ───────┤
       |                               |
       └── Security Best Practices ────┤
                                       |
                                       ▼
                        Production-Ready Applications
```

- **Architecture**: Scalable and maintainable code structure
- **Testing**: Comprehensive quality assurance strategies
- **Authentication**: Secure user identity management
- **Security**: Protection against common vulnerabilities

</div>

---

# Part I: Clean Architecture

<div class="slide-content">

> **Clean Architecture** is a software design philosophy that separates concerns and makes applications independent of frameworks, databases, and external agencies.

#### Core principles
- **Independence**: Business logic doesn't depend on external details
- **Testability**: Each layer can be tested in isolation
- **Flexibility**: Easy to change frameworks, databases, or UI
- **Maintainability**: Clear boundaries and responsibilities

#### Benefits
- **Reduced coupling**: Changes in one layer don't affect others
- **Improved testability**: Mock external dependencies easily
- **Better team collaboration**: Clear boundaries for different teams
- **Future-proofing**: Easy to adapt to new requirements

</div>

---

# Clean Architecture layers

<div class="slide-content">

<div class="code-columns">
<div>

#### Traditional layered architecture
```
┌─────────────────────┐
│    Presentation     │ ← Controllers, Views
├─────────────────────┤
│      Business       │ ← Services, Logic
├─────────────────────┤
│    Data Access      │ ← Repositories
├─────────────────────┤
│      Database       │ ← PostgreSQL, Files
└─────────────────────┘
```

**Problems:**
- Database drives design
- Business logic depends on data layer
- Hard to test business rules
- Framework lock-in

</div>

<div>

#### Clean architecture layers
```
┌─────────────────────┐
│    Frameworks       │ ← Web, CLI, Mobile
├─────────────────────┤
│   Interface/UI      │ ← Controllers, Views
├─────────────────────┤
│   Use Cases         │ ← Application Logic
├─────────────────────┤
│     Entities        │ ← Business Rules
└─────────────────────┘
      ↑ Dependencies point inward
```

**Benefits:**
- Business rules at center
- Framework independence
- Testable in isolation
- Dependency inversion

</div>

</div>

</div>

---

# Go clean architecture structure

<div class="slide-content">

<div class="code-columns">
<div>

#### Project structure
```
project/
├── cmd/
│   └── server/          # Main entry point
├── internal/
│   ├── domain/          # Entities & interfaces
│   ├── usecase/         # Application logic
│   ├── repository/      # Data access
│   ├── handler/         # HTTP handlers
│   └── middleware/      # Cross-cutting concerns
├── pkg/                 # Shared utilities
└── migrations/          # Database schemas
```

</div>

<div>

#### Dependency direction
```go
// domain/user.go - Core entities
type User struct {
    ID    int
    Name  string
    Email string
}

// domain/interfaces.go - Abstractions
type UserRepository interface {
    Save(user User) error
    FindByID(id int) (*User, error)
}

type UserUseCase interface {
    CreateUser(name, email string) (*User, error)
    GetUser(id int) (*User, error)
}
```

</div>

</div>

</div>

---

# Domain layer implementation

<div class="slide-content">

<div class="code-columns">
<div>

#### Entities (business objects)
```go
// domain/user.go
type User struct {
    ID        int
    Name      string
    Email     string
    CreatedAt time.Time
}

func (u *User) Validate() error {
    if u.Name == "" {
        return errors.New("name required")
    }
    if !isValidEmail(u.Email) {
        return errors.New("invalid email")
    }
    return nil
}
```

</div>

<div>

#### Repository interfaces
```go
// domain/repository.go
type UserRepository interface {
    Save(user *User) error
    FindByID(id int) (*User, error)
    FindByEmail(email string) (*User, error)
    Update(user *User) error
    Delete(id int) error
}

type PostRepository interface {
    Save(post *Post) error
    FindByUserID(userID int) ([]*Post, error)
    FindPublished() ([]*Post, error)
}
```

</div>

</div>

</div>

---

# Use case layer implementation

<div class="slide-content">

<div class="code-columns">
<div>

#### User use cases
```go
// usecase/user.go
type UserUseCase struct {
    userRepo UserRepository
    logger   Logger
}

func (uc *UserUseCase) CreateUser(name, email string) (*User, error) {
    // Validate input
    user := &User{Name: name, Email: email}
    if err := user.Validate(); err != nil {
        return nil, err
    }
    
    // Check if user exists
    existing, _ := uc.userRepo.FindByEmail(email)
    if existing != nil {
        return nil, errors.New("user already exists")
    }
    
    // Save user
    if err := uc.userRepo.Save(user); err != nil {
        uc.logger.Error("Failed to save user", err)
        return nil, err
    }
    
    return user, nil
}
```

</div>

<div>

#### Dependency injection
```go
// usecase/interfaces.go
type UserUseCase interface {
    CreateUser(name, email string) (*User, error)
    GetUser(id int) (*User, error)
    UpdateUser(id int, name, email string) error
    DeleteUser(id int) error
}

// Wire dependencies
func NewUserUseCase(repo UserRepository, logger Logger) UserUseCase {
    return &userUseCase{
        userRepo: repo,
        logger:   logger,
    }
}
```

</div>

</div>

</div>

---

# Repository implementation

<div class="slide-content">

<div class="code-columns">
<div>

#### PostgreSQL repository
```go
// repository/user_postgres.go
type postgresUserRepository struct {
    db *sql.DB
}

func (r *postgresUserRepository) Save(user *User) error {
    query := `INSERT INTO users (name, email) 
              VALUES ($1, $2) RETURNING id, created_at`
    
    return r.db.QueryRow(query, user.Name, user.Email).
        Scan(&user.ID, &user.CreatedAt)
}

func (r *postgresUserRepository) FindByID(id int) (*User, error) {
    query := `SELECT id, name, email, created_at 
              FROM users WHERE id = $1`
    
    user := &User{}
    err := r.db.QueryRow(query, id).
        Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
    
    return user, err
}
```

</div>

<div>

#### Memory repository (for testing)
```go
// repository/user_memory.go
type memoryUserRepository struct {
    users map[int]*User
    mutex sync.RWMutex
    nextID int
}

func (r *memoryUserRepository) Save(user *User) error {
    r.mutex.Lock()
    defer r.mutex.Unlock()
    
    r.nextID++
    user.ID = r.nextID
    user.CreatedAt = time.Now()
    r.users[user.ID] = user
    
    return nil
}

func (r *memoryUserRepository) FindByID(id int) (*User, error) {
    r.mutex.RLock()
    defer r.mutex.RUnlock()
    
    user, exists := r.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}
```

</div>

</div>

</div>

---

# Flutter clean architecture

<div class="slide-content">

<div class="code-columns">
<div>

#### Project structure
```
lib/
├── core/
│   ├── error/           # Error handling
│   ├── network/         # Network utilities
│   └── usecases/        # Base use case
├── features/
│   └── user/
│       ├── data/        # Data sources & repos
│       ├── domain/      # Entities & use cases
│       └── presentation/# UI & state management
└── injection/           # Dependency injection
```

</div>

<div>

#### Entities
```dart
// features/user/domain/entities/user.dart
class User extends Equatable {
  final int id;
  final String name;
  final String email;
  final DateTime createdAt;

  const User({
    required this.id,
    required this.name,
    required this.email,
    required this.createdAt,
  });

  @override
  List<Object> get props => [id, name, email, createdAt];
}
```

</div>

</div>

</div>

---

# Flutter use cases

<div class="slide-content">

<div class="code-columns">
<div>

#### Base use case
```dart
// core/usecases/usecase.dart
abstract class UseCase<Type, Params> {
  Future<Either<Failure, Type>> call(Params params);
}

class NoParams extends Equatable {
  @override
  List<Object> get props => [];
}
```

#### Specific use case
```dart
// features/user/domain/usecases/get_user.dart
class GetUser implements UseCase<User, GetUserParams> {
  final UserRepository repository;

  GetUser(this.repository);

  @override
  Future<Either<Failure, User>> call(GetUserParams params) async {
    return await repository.getUser(params.id);
  }
}
```

</div>

<div>

#### Repository contract
```dart
// features/user/domain/repositories/user_repository.dart
abstract class UserRepository {
  Future<Either<Failure, User>> getUser(int id);
  Future<Either<Failure, User>> createUser(String name, String email);
  Future<Either<Failure, List<User>>> getUsers();
}

// Parameters class
class GetUserParams extends Equatable {
  final int id;

  const GetUserParams({required this.id});

  @override
  List<Object> get props => [id];
}
```

</div>

</div>

</div>

---

# Data layer implementation

<div class="slide-content">

<div class="code-columns">
<div>

#### Data models
```dart
// features/user/data/models/user_model.dart
class UserModel extends User {
  const UserModel({
    required super.id,
    required super.name,
    required super.email,
    required super.createdAt,
  });

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      id: json['id'],
      name: json['name'],
      email: json['email'],
      createdAt: DateTime.parse(json['created_at']),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'email': email,
      'created_at': createdAt.toIso8601String(),
    };
  }
}
```

</div>

<div>

#### Repository implementation
```dart
// features/user/data/repositories/user_repository_impl.dart
class UserRepositoryImpl implements UserRepository {
  final UserRemoteDataSource remoteDataSource;
  final UserLocalDataSource localDataSource;
  final NetworkInfo networkInfo;

  UserRepositoryImpl({
    required this.remoteDataSource,
    required this.localDataSource,
    required this.networkInfo,
  });

  @override
  Future<Either<Failure, User>> getUser(int id) async {
    if (await networkInfo.isConnected) {
      try {
        final user = await remoteDataSource.getUser(id);
        await localDataSource.cacheUser(user);
        return Right(user);
      } catch (e) {
        return Left(ServerFailure());
      }
    } else {
      try {
        final user = await localDataSource.getUser(id);
        return Right(user);
      } catch (e) {
        return Left(CacheFailure());
      }
    }
  }
}
```

</div>

</div>

</div>

---

# Part II: Testing Strategies

<div class="slide-content">

> **Testing** is crucial for building reliable applications. Different types of tests serve different purposes and provide varying levels of confidence.

#### Testing pyramid
```
                    ▲
                   /|\
                  / | \
                 /  |  \
                / E2E|   \
               /Tests|    \
              /______|_____\
             /              \
            /  Integration    \
           /     Tests        \
          /___________________\
         /                     \
        /      Unit Tests       \
       /_______________________\
```

- **Unit Tests**: Fast, isolated, test individual functions
- **Integration Tests**: Test component interactions
- **End-to-End Tests**: Test complete user workflows

</div>

---

# Go testing fundamentals

<div class="slide-content">

<div class="code-columns">
<div>

#### Basic unit test
```go
// user_test.go
func TestUser_Validate(t *testing.T) {
    tests := []struct {
        name    string
        user    User
        wantErr bool
    }{
        {"valid user", User{Name: "John", Email: "john@test.com"}, false},
        {"empty name", User{Name: "", Email: "john@test.com"}, true},
        {"invalid email", User{Name: "John", Email: "invalid"}, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.user.Validate()
            if (err != nil) != tt.wantErr {
                t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

</div>

<div>

#### Testing with mocks
```go
// usecase_test.go
func TestUserUseCase_CreateUser(t *testing.T) {
    mockRepo := &MockUserRepository{}
    mockLogger := &MockLogger{}
    useCase := NewUserUseCase(mockRepo, mockLogger)
    
    // Setup mock expectations
    mockRepo.On("FindByEmail", "test@example.com").Return(nil, nil)
    mockRepo.On("Save", mock.AnythingOfType("*User")).Return(nil)
    
    // Execute
    user, err := useCase.CreateUser("Test User", "test@example.com")
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "Test User", user.Name)
    mockRepo.AssertExpectations(t)
}
```

</div>

</div>

</div>

---

# Testing repository layer

<div class="slide-content">

<div class="code-columns">
<div>

#### Integration test with database
```go
// repository_test.go
func TestPostgresUserRepository_Save(t *testing.T) {
    db := setupTestDB(t)
    defer cleanupTestDB(t, db)
    
    repo := NewPostgresUserRepository(db)
    user := &User{
        Name:  "Test User",
        Email: "test@example.com",
    }
    
    err := repo.Save(user)
    
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)
    assert.NotZero(t, user.CreatedAt)
}
```

</div>

<div>

#### Test helpers
```go
// test_helpers.go
func setupTestDB(t *testing.T) *sql.DB {
    db, err := sql.Open("postgres", "postgres://test:test@localhost/testdb?sslmode=disable")
    require.NoError(t, err)
    
    // Run migrations
    err = runMigrations(db)
    require.NoError(t, err)
    
    return db
}

func cleanupTestDB(t *testing.T, db *sql.DB) {
    _, err := db.Exec("TRUNCATE users, posts RESTART IDENTITY CASCADE")
    require.NoError(t, err)
    db.Close()
}
```

</div>

</div>

</div>

---

# Flutter testing approaches

<div class="slide-content">

<div class="code-columns">
<div>

#### Unit testing entities
```dart
// test/features/user/domain/entities/user_test.dart
void main() {
  group('User', () {
    test('should be a subclass of Equatable', () {
      // arrange
      const user = User(id: 1, name: 'Test', email: 'test@test.com', createdAt: DateTime(2025));
      
      // assert
      expect(user, isA<Equatable>());
    });

    test('should return correct props', () {
      // arrange
      const user = User(id: 1, name: 'Test', email: 'test@test.com', createdAt: DateTime(2025));
      
      // assert
      expect(user.props, [1, 'Test', 'test@test.com', DateTime(2025)]);
    });
  });
}
```

</div>

<div>

#### Testing use cases
```dart
// test/features/user/domain/usecases/get_user_test.dart
void main() {
  late GetUser usecase;
  late MockUserRepository mockRepository;

  setUp(() {
    mockRepository = MockUserRepository();
    usecase = GetUser(mockRepository);
  });

  test('should get user from repository', () async {
    // arrange
    const testUser = User(id: 1, name: 'Test', email: 'test@test.com', createdAt: DateTime(2025));
    when(mockRepository.getUser(any)).thenAnswer((_) async => const Right(testUser));

    // act
    final result = await usecase(const GetUserParams(id: 1));

    // assert
    expect(result, const Right(testUser));
    verify(mockRepository.getUser(1));
    verifyNoMoreInteractions(mockRepository);
  });
}
```

</div>

</div>

</div>

---

# Widget testing

<div class="slide-content">

<div class="code-columns">
<div>

#### Basic widget test
```dart
// test/features/user/presentation/widgets/user_card_test.dart
void main() {
  testWidgets('UserCard displays user information', (tester) async {
    // arrange
    const user = User(
      id: 1,
      name: 'John Doe',
      email: 'john@example.com',
      createdAt: DateTime(2025),
    );

    // act
    await tester.pumpWidget(
      MaterialApp(
        home: Scaffold(
          body: UserCard(user: user),
        ),
      ),
    );

    // assert
    expect(find.text('John Doe'), findsOneWidget);
    expect(find.text('john@example.com'), findsOneWidget);
  });
}
```

</div>

<div>

#### Testing with BLoC
```dart
// test/features/user/presentation/bloc/user_bloc_test.dart
void main() {
  late UserBloc bloc;
  late MockGetUser mockGetUser;

  setUp(() {
    mockGetUser = MockGetUser();
    bloc = UserBloc(getUser: mockGetUser);
  });

  group('GetUserEvent', () {
    test('should emit [Loading, Loaded] when data is gotten successfully', () {
      // arrange
      when(mockGetUser(any)).thenAnswer((_) async => const Right(testUser));

      // assert later
      final expected = [
        UserLoading(),
        const UserLoaded(user: testUser),
      ];
      expectLater(bloc.stream, emitsInOrder(expected));

      // act
      bloc.add(const GetUserEvent(id: 1));
    });
  });
}
```

</div>

</div>

</div>

---

# Test coverage and quality

<div class="slide-content">

<div class="code-columns">
<div>

#### Go test coverage
```bash
# Run tests with coverage
go test -v -cover ./...

# Generate detailed coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Set coverage threshold
go test -cover ./... | grep "coverage:"
```

#### Coverage goals
- **Domain layer**: 95%+ (business logic)
- **Use cases**: 90%+ (application logic)
- **Repositories**: 80%+ (data access)
- **Handlers**: 70%+ (HTTP layer)

</div>

<div>

#### Flutter test coverage
```bash
# Run tests with coverage
flutter test --coverage

# Generate HTML report
genhtml coverage/lcov.info -o coverage/html

# View coverage
open coverage/html/index.html
```

#### Best practices
- **Test behavior, not implementation**
- **One assertion per test**
- **Use descriptive test names**
- **Setup and teardown properly**
- **Mock external dependencies**

</div>

</div>

</div>

---

# Part III: JWT & Authentication

<div class="slide-content">

> **JWT (JSON Web Tokens)** provide a stateless way to handle authentication and authorization in distributed systems.

#### JWT structure
```
Header.Payload.Signature
```

#### Components
- **Header**: Algorithm and token type
- **Payload**: Claims (user data, permissions)
- **Signature**: Cryptographic signature for verification

#### Benefits
- **Stateless**: No server-side session storage
- **Scalable**: Works across multiple services
- **Portable**: Can include user claims
- **Secure**: Cryptographically signed

</div>

---

# JWT implementation in Go

<div class="slide-content">

<div class="code-columns">
<div>

#### JWT service
```go
// internal/auth/jwt.go
type JWTService struct {
    secretKey string
    issuer    string
    expiry    time.Duration
}

func (j *JWTService) GenerateToken(userID int, email string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "email":   email,
        "iss":     j.issuer,
        "exp":     time.Now().Add(j.expiry).Unix(),
        "iat":     time.Now().Unix(),
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(j.secretKey))
}
```

</div>

<div>

#### Token validation
```go
func (j *JWTService) ValidateToken(tokenString string) (*Claims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return []byte(j.secretKey), nil
    })
    
    if err != nil || !token.Valid {
        return nil, errors.New("invalid token")
    }
    
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, errors.New("invalid claims")
    }
    
    return &Claims{
        UserID: int(claims["user_id"].(float64)),
        Email:  claims["email"].(string),
    }, nil
}
```

</div>

</div>

</div>

---

# Authentication middleware

<div class="slide-content">

<div class="code-columns">
<div>

#### JWT middleware
```go
// internal/middleware/auth.go
func JWTAuthMiddleware(jwtService *JWTService) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                http.Error(w, "Authorization header required", http.StatusUnauthorized)
                return
            }
            
            tokenString := strings.TrimPrefix(authHeader, "Bearer ")
            claims, err := jwtService.ValidateToken(tokenString)
            if err != nil {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }
            
            ctx := context.WithValue(r.Context(), "user", claims)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

</div>

<div>

#### Protected routes
```go
// internal/handlers/protected.go
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
    claims, ok := r.Context().Value("user").(*Claims)
    if !ok {
        http.Error(w, "User not found in context", http.StatusInternalServerError)
        return
    }
    
    user, err := userService.GetUser(claims.UserID)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    writeJSON(w, http.StatusOK, user)
}

// Apply middleware
r.Handle("/api/profile", JWTAuthMiddleware(jwtService)(http.HandlerFunc(GetUserProfile)))
```

</div>

</div>

</div>

---

# Flutter JWT handling

<div class="slide-content">

<div class="code-columns">
<div>

#### Token storage
```dart
// core/auth/token_storage.dart
class TokenStorage {
  static const _tokenKey = 'auth_token';
  static const _refreshTokenKey = 'refresh_token';
  
  static Future<void> saveTokens(String token, String refreshToken) async {
    await SecureStorage.write(_tokenKey, token);
    await SecureStorage.write(_refreshTokenKey, refreshToken);
  }
  
  static Future<String?> getToken() async {
    return await SecureStorage.read(_tokenKey);
  }
  
  static Future<String?> getRefreshToken() async {
    return await SecureStorage.read(_refreshTokenKey);
  }
  
  static Future<void> clearTokens() async {
    await SecureStorage.delete(_tokenKey);
    await SecureStorage.delete(_refreshTokenKey);
  }
}
```

</div>

<div>

#### Auth interceptor
```dart
// core/network/auth_interceptor.dart
class AuthInterceptor extends Interceptor {
  @override
  void onRequest(RequestOptions options, RequestInterceptorHandler handler) async {
    final token = await TokenStorage.getToken();
    if (token != null) {
      options.headers['Authorization'] = 'Bearer $token';
    }
    super.onRequest(options, handler);
  }

  @override
  void onError(DioError err, ErrorInterceptorHandler handler) async {
    if (err.response?.statusCode == 401) {
      final refreshed = await _refreshToken();
      if (refreshed) {
        // Retry request with new token
        final newToken = await TokenStorage.getToken();
        err.requestOptions.headers['Authorization'] = 'Bearer $newToken';
        final response = await Dio().fetch(err.requestOptions);
        return handler.resolve(response);
      }
    }
    super.onError(err, handler);
  }
}
```

</div>

</div>

</div>

---

# Part IV: Security Best Practices

<div class="slide-content">

> **Security** is not an afterthought but should be built into every layer of your application from the beginning.

#### Common vulnerabilities (OWASP Top 10)
- **Injection**: SQL, NoSQL, Command injection
- **Broken Authentication**: Weak passwords, session management
- **Sensitive Data Exposure**: Unencrypted data transmission/storage
- **XML External Entities (XXE)**: Processing untrusted XML
- **Broken Access Control**: Unauthorized access to resources
- **Security Misconfiguration**: Default configs, unnecessary features
- **Cross-Site Scripting (XSS)**: Malicious script injection
- **Insecure Deserialization**: Arbitrary code execution
- **Using Components with Known Vulnerabilities**: Outdated dependencies

</div>

---

# Input validation and sanitization

<div class="slide-content">

<div class="code-columns">
<div>

#### Go input validation
```go
// internal/validators/user.go
type CreateUserRequest struct {
    Name     string `json:"name" validate:"required,min=2,max=100"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}

func ValidateCreateUserRequest(req *CreateUserRequest) error {
    validate := validator.New()
    if err := validate.Struct(req); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }
    
    // Custom business rules
    if containsBadWords(req.Name) {
        return errors.New("name contains inappropriate content")
    }
    
    return nil
}
```

</div>

<div>

#### SQL injection prevention
```go
// ❌ Vulnerable to SQL injection
query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", email)

// ✅ Safe with prepared statements
query := "SELECT * FROM users WHERE email = $1"
rows, err := db.Query(query, email)

// ✅ Safe with query builder
query := squirrel.Select("*").From("users").Where(squirrel.Eq{"email": email})
sql, args, _ := query.ToSql()
rows, err := db.Query(sql, args...)
```

</div>

</div>

</div>

---

# Password security

<div class="slide-content">

<div class="code-columns">
<div>

#### Password hashing
```go
// internal/auth/password.go
import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
    // Use bcrypt with cost 12 (recommended)
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    return string(bytes), err
}

func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// Password strength validation
func ValidatePasswordStrength(password string) error {
    if len(password) < 8 {
        return errors.New("password must be at least 8 characters")
    }
    // Add more complexity rules
    return nil
}
```

</div>

<div>

#### Rate limiting
```go
// internal/middleware/rate_limit.go
import "golang.org/x/time/rate"

func RateLimitMiddleware(rps rate.Limit, burst int) func(http.Handler) http.Handler {
    limiter := rate.NewLimiter(rps, burst)
    
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if !limiter.Allow() {
                http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}

// Apply to login endpoint
r.Handle("/api/login", RateLimitMiddleware(rate.Every(time.Second), 5)(loginHandler))
```

</div>

</div>

</div>

---

# HTTPS and TLS configuration

<div class="slide-content">

<div class="code-columns">
<div>

#### TLS server configuration
```go
// cmd/server/main.go
func setupTLSServer() *http.Server {
    server := &http.Server{
        Addr:         ":8443",
        Handler:      setupRoutes(),
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
        TLSConfig: &tls.Config{
            MinVersion:         tls.VersionTLS12,
            CurvePreferences:   []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
            PreferServerCiphers: true,
            CipherSuites: []uint16{
                tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
                tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
                tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
            },
        },
    }
    return server
}
```

</div>

<div>

#### Security headers
```go
// internal/middleware/security.go
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // HTTPS enforcement
        w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
        
        // XSS Protection
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        
        // Content Security Policy
        w.Header().Set("Content-Security-Policy", "default-src 'self'")
        
        next.ServeHTTP(w, r)
    })
}
```

</div>

</div>

</div>

---

# Flutter security considerations

<div class="slide-content">

<div class="code-columns">
<div>

#### Certificate pinning
```dart
// core/network/certificate_pinning.dart
class CertificatePinner {
  static Dio createPinnedClient() {
    final dio = Dio();
    
    // Add certificate pinning
    (dio.httpClientAdapter as DefaultHttpClientAdapter).onHttpClientCreate = (client) {
      client.badCertificateCallback = (cert, host, port) {
        // Verify certificate fingerprint
        final fingerprint = sha256.convert(cert.der).toString();
        return fingerprint == expectedFingerprint;
      };
      return client;
    };
    
    return dio;
  }
}
```

</div>

<div>

#### Secure data handling
```dart
// core/security/data_protection.dart
class DataProtection {
  // Encrypt sensitive data before storage
  static Future<String> encryptData(String data) async {
    final key = await _getOrCreateKey();
    final encrypted = AES(key).encrypt(data);
    return encrypted.base64;
  }
  
  // Sanitize sensitive data from logs
  static Map<String, dynamic> sanitizeForLogging(Map<String, dynamic> data) {
    final sanitized = Map<String, dynamic>.from(data);
    
    // Remove sensitive fields
    sanitized.remove('password');
    sanitized.remove('token');
    sanitized.remove('credit_card');
    
    return sanitized;
  }
}
```

</div>

</div>

</div>

---

# Security checklist

<div class="slide-content">

#### Backend Security ✅
- [ ] **Input validation** on all endpoints
- [ ] **SQL injection** prevention with prepared statements
- [ ] **Authentication** with secure JWT implementation
- [ ] **Rate limiting** on sensitive endpoints
- [ ] **HTTPS** with proper TLS configuration
- [ ] **Security headers** (HSTS, CSP, X-Frame-Options)
- [ ] **Password hashing** with bcrypt or similar
- [ ] **Environment variables** for secrets
- [ ] **Dependency scanning** for vulnerabilities
- [ ] **Error handling** without information leakage

#### Frontend Security ✅
- [ ] **Certificate pinning** for API connections
- [ ] **Secure storage** for sensitive data
- [ ] **Input sanitization** for user inputs
- [ ] **Token management** with automatic refresh
- [ ] **Deep link validation** and authorization
- [ ] **Biometric authentication** where appropriate
- [ ] **App transport security** configuration
- [ ] **Code obfuscation** for production builds

</div>

---

# What we've learned

<div class="slide-content">

## Clean Architecture
- **Separation of concerns**: Domain, use case, and infrastructure layers
- **Dependency inversion**: Abstract interfaces over concrete implementations
- **Testability**: Each layer can be tested in isolation
- **Maintainability**: Clear boundaries and responsibilities

## Testing Strategies
- **Testing pyramid**: Unit, integration, and end-to-end tests
- **Mocking**: Isolating units under test from dependencies
- **Coverage**: Measuring and improving test coverage
- **Quality**: Best practices for reliable and maintainable tests

</div>

---

# What we've learned (continued)

<div class="slide-content">

## JWT & Authentication
- **Stateless authentication**: JWT structure and benefits
- **Token management**: Generation, validation, and refresh
- **Security considerations**: Secret management and token expiry
- **Implementation**: Go JWT services and Flutter token handling

## Security Best Practices
- **Input validation**: Preventing injection attacks
- **Password security**: Proper hashing and strength requirements
- **HTTPS/TLS**: Secure communication channels
- **Security headers**: Browser protection mechanisms
- **Mobile security**: Certificate pinning and secure storage

</div>

---

# Thank You!

<div class="slide-content">

**What's Next:**
- Lab 05: Implement clean architecture with comprehensive testing and security

**Resources:**
- Clean Architecture: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
- Go Testing: https://pkg.go.dev/testing
- Flutter Testing: https://docs.flutter.dev/testing
- JWT Best Practices: https://auth0.com/blog/a-look-at-the-latest-draft-for-jwt-bcp/
- OWASP Top 10: https://owasp.org/www-project-top-ten/
- Course Repository: https://github.com/timur-harin/sum25-go-flutter-course

**Contact:**
- Email: timur.harin@mail.com
- Telegram: @timur_harin

**Next Lecture:** Deployment & DevOps

## Questions?

</div> 