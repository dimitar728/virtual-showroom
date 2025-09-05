# React / Golang - User Management Application

This is a full-stack user management application built with React, Redux, TypeScript, and Tailwind CSS for the frontend, and Go with PostgreSQL for the backend.



![image](https://github.com/user-attachments/assets/02812736-08d1-4fa4-8570-5d2e03a8dfb1)


![image](https://github.com/user-attachments/assets/e69f5332-c40b-43b5-9003-10920bf34ef5)


## Frontend (React, Redux, TypeScript, Tailwind CSS)

### Setup

1. Navigate to the frontend directory:
   ```
   cd frontend
   ```

2. Install dependencies:
   ```
   npm install
   ```

3. Start the development server:
   ```
   npm start
   ```

The application will be available at `http://localhost:3000`.

### Features

- User registration with name, email, and password
- User login with email and password
- JWT-based authentication
- Protected routes for authenticated users
- Responsive design using Tailwind CSS
- State management with Redux Toolkit
- Type-safe development with TypeScript

### Key Components

- `Register`: Allows new users to create an account
- `Login`: Allows existing users to log in
- `UserList`: Displays a list of users (accessible to authenticated users)
- `UserForm`: Allows editing user information
- `Navbar`: Navigation component with dynamic links based on authentication status

### State Management

- Uses Redux Toolkit for state management
- `authSlice`: Manages authentication state, including user info and token
- `userSlice`: Manages user data for the user list and user editing functionality

## Backend (Go, PostgreSQL)

### Setup

1. Navigate to the backend directory:
   ```
   cd backend
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Set up your PostgreSQL database and update the `.env` file with your database credentials.

4. Run the application:
   ```
   go run cmd/main.go
   ```

The server will start on `http://localhost:8080`.

### Features

- RESTful API endpoints for user management
- JWT-based authentication
- PostgreSQL database for data persistence
- GORM as the ORM for database operations
- Gin web framework for routing and middleware

### Key Endpoints

- `POST /api/register`: Register a new user
- `POST /api/login`: Authenticate a user and return a JWT
- `GET /api/users`: Retrieve a list of users (protected route)
- `GET /api/users/:id`: Retrieve a specific user (protected route)
- `PUT /api/users/:id`: Update a user's information (protected route)
- `DELETE /api/users/:id`: Delete a user (protected route)

### Database Schema

The `users` table includes the following fields:
- `id`: Unique identifier (auto-incremented)
- `name`: User's full name
- `email`: User's email address (unique)
- `password`: Hashed password
- `created_at`: Timestamp of user creation
- `updated_at`: Timestamp of last update

## Technologies Used

### Frontend
- React
- Redux Toolkit
- TypeScript
- Tailwind CSS
- Axios for API requests
- React Router for navigation

### Backend
- Go
- Gin web framework
- GORM (Go Object Relational Mapper)
- PostgreSQL
- JWT for authentication

## Getting Started

1. Clone the repository
2. Set up the backend:
   - Install Go and PostgreSQL
   - Set up your database and update the `.env` file
   - Run the Go application
3. Set up the frontend:
   - Install Node.js and npm
   - Install frontend dependencies
   - Start the React development server
4. Access the application at `http://localhost:3000`


## PostgreSQL Database Management

To manage your PostgreSQL database directly, you can use the following commands in the psql interactive terminal.

### Connecting to the Database

Connect to your PostgreSQL database using:

```
psql -U alan_jones -d user_management
```

You'll be prompted for the password you set for the `alan_jones` user.

### Useful PostgreSQL Commands

Once connected, you can use these SQL commands to manage and view your data:

```sql
-- View all users
SELECT * FROM users;

-- View specific columns
SELECT id, name, email, created_at FROM users;

-- Search for a specific user by email
SELECT * FROM users WHERE email = 'user@example.com';

-- Count the number of users
SELECT COUNT(*) FROM users;

-- View the most recently registered users (limit to 5)
SELECT * FROM users ORDER BY created_at DESC LIMIT 5;

-- View users registered in the last 24 hours
SELECT * FROM users WHERE created_at > NOW() - INTERVAL '1 day';

-- Update a user's name
UPDATE users SET name = 'New Name' WHERE id = 1;

-- Delete a user (be careful with this!)
DELETE FROM users WHERE id = 1;

-- View the table structure
\d users
```

### Exiting psql

To exit the psql prompt, type:

```
\q
```

### Database Maintenance

If you need to perform database maintenance or ensure all connections are closed:

On Unix-like systems (Linux, macOS):
```
sudo service postgresql restart
```
or
```
sudo systemctl restart postgresql
```

On Windows (in an administrator command prompt):
```
net stop postgresql && net start postgresql
```

Remember to replace `user@example.com`, `'New Name'`, and the ID numbers with actual values when using these commands. Always exercise caution when running UPDATE or DELETE commands, especially in a production environment.

## Future Enhancements

- Implement password reset functionality
- Add user roles and permissions
- Implement email verification
- Add unit and integration tests
- Set up CI/CD pipeline


