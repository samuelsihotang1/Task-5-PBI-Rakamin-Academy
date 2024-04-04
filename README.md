#### My Rakamin Profile : https://www.rakamin.com/profile/samuel-christy-angie-sihotang-e2mi2sqfttox765s

# Task 5 PBI - Web Application

## Overview
This repository contains the source code for a web application developed for Task 5 PBI project. The application serves as a platform where users can register, login, upload photos, edit their profiles, and perform other related tasks. Below is a breakdown of the main components and functionalities of the application:

### File Structure
- **Connect.go:** Establishes a connection to the MySQL database.
- **Migrate.go:** Handles the database migrations, ensuring the necessary tables are created.
- **Model.go:** Defines a generic model struct with common fields like ID, CreatedAt, and UpdatedAt.
- **User.go:** Contains the model definition for the User entity, including fields such as Username, Email, Password, and a reference to Photos.
- **Photo.go:** Defines the model for Photo entities, including Title, Caption, PhotoUrl, and a reference to the User who uploaded it.
- **RequireAuth.go:** Middleware function to ensure authentication for protected routes.
- **LoadEnvVariable.go:** Helper function to load environment variables from a .env file.
- **ValidateStruct.go:** Helper function for validating request body structures.
- **controllers/:** Contains handlers for various HTTP requests.
  - **UserController.go:** Handles user-related operations such as registration, login, editing, and deletion.
  - **PhotoController.go:** Manages photo-related operations like uploading, retrieving, editing, and deletion.
- **router/:** Defines HTTP routes and connects them to the corresponding controllers.
  - **routers.go:** Specifies the routes for user authentication, profile management, and photo-related actions.

### Routing
- **/users/register (POST):** Allows users to register by providing a username, email, and password. Upon successful registration, a JWT token is generated, and the user is automatically logged in.
- **/register (GET):** Renders the registration page.
- **/users/login (POST):** Handles user login. Users provide their email and password, and upon successful authentication, a JWT token is generated and stored as a cookie for subsequent requests.
- **/login (GET):** Renders the login page.
- **/ (GET):** Serves the homepage after successful authentication.
- **/users/info (GET):** Retrieves user information including username and email.
- **/users (GET):** Renders a page for user profile management.
- **/users/:userId (PUT):** Allows users to edit their profile information such as username, email, and password.
- **/users/logout (POST):** Logs the user out by deleting the JWT token cookie.
- **/users/:userId (DELETE):** Allows users to delete their account.
- **/photos (POST):** Enables users to upload photos along with a title and caption.
- **/photos (GET):** Retrieves the first photo uploaded by the current user.
- **/infophoto (GET):** Retrieves information about the first photo uploaded by the current user.
- **/photos/:photoId (PUT):** Allows users to edit the title and caption of a specific photo.
- **/photos/:photoId (DELETE):** Enables users to delete a specific photo they uploaded.

### Functionality
- **User Management:** Users can register, login, edit their profile information, and delete their accounts.
- **Authentication:** JWT-based authentication ensures secure access to protected routes.
- **Photo Management:** Users can upload photos, view their uploaded photos, edit photo details, and delete photos they uploaded.
- **Middleware:** Middleware functions ensure authentication for protected routes, preventing unauthorized access.

### How It Works
1. Users can register for an account by providing a username, email, and password.
2. Upon successful registration, users are automatically logged in, and a JWT token is stored as a cookie.
3. Users can log in using their email and password, which validates their credentials against the database.
4. After logging in, users can access protected routes such as their profile page and photo management.
5. Users can upload photos with titles and captions, view their uploaded photos, edit photo details, and delete photos.
6. User profile information can be edited, including username, email, and password.
7. Users can log out, which deletes the JWT token cookie, terminating their session.

### Setup Instructions
1. Clone the repository to your local machine.
2. Create a `.env` file and define the necessary environment variables like `DB`, `SECRET`, and `PORT`.
3. Run `go run main.go` to start the application.
4. Access the application via the specified routes, e.g., `http://localhost:3000/login` for the login page.

## Contributors
- Samuel Christy Angie Sihotang
