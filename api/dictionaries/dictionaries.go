package dictionaries

// Authentication messages
const (
	// Success
	UserRegisteredSuccess = "User registered successfully"
	UserLoggedSuccess     = "User logged in successfully"
	MasterCreatedSuccess  = "Master user created successfully"
	MasterDeletedSuccess  = "Master user deleted successfully"

	// Error
	InvalidCredentials    = "Invalid credentials"
	UserAlreadyExists     = "User already exists"
	UserNotFound          = "User not found"
	InvalidToken          = "Invalid token"
	UnauthorizedAccess    = "Unauthorized access"
	MasterAlreadyExists   = "Master user already exists"
	MasterNotFound        = "Master user not found"
	InvalidMasterPassword = "Invalid master password"
)

// User messages
const (
	// Success
	UserCreatedSuccess = "User created successfully"
	UserDeletedSuccess = "User deleted successfully"

	// Error
	UserCreationFailed = "Failed to create user"
	UserDeletionFailed = "Failed to delete user"
	UserNotAuthorized  = "User not authorized for this operation"
	InvalidUserRole    = "Invalid user role"
)

// Ticket messages
const (
	// Success
	TicketCreatedSuccess = "Ticket created successfully"
	TicketDeletedSuccess = "Ticket deleted successfully"
	TicketStatusUpdated  = "Ticket status updated successfully"
	TicketHistoryAdded   = "Ticket history updated successfully"
	TicketFoundSuccess   = "Ticket found successfully"
	TicketsListedSuccess = "Tickets listed successfully"

	// Error
	TicketCreationFailed     = "Failed to create ticket"
	TicketDeletionFailed     = "Failed to delete ticket"
	TicketNotFound           = "Ticket not found"
	TicketStatusUpdateFailed = "Failed to update ticket status"
	TicketHistoryAddFailed   = "Failed to add ticket history"
	NoTicketsFound           = "No tickets found"
	NoTicketsForAuthor       = "Author has no tickets"
	NoTicketsForStatus       = "No tickets found with specified status"
	InvalidTicketStatus      = "Invalid ticket status"
	InvalidTicketData        = "Invalid ticket data"
	NoPermissionToDelete     = "You don't have permission to delete this ticket"
	InvalidDateFormat        = "Invalid date format"
)

// Image messages
const (
	// Success
	ImageUploadedSuccess = "Image uploaded successfully"

	// Error
	ImageUploadFailed   = "Failed to upload image"
	InvalidImageFormat  = "Invalid image format"
	InvalidImageContent = "Invalid image content"
	ImageTooLarge       = "Image too large"
)

// General messages
const (
	// Success
	OperationSuccess = "Operation completed successfully"

	// Error
	InvalidData          = "Invalid data"
	InternalServerError  = "Internal server error"
	DatabaseError        = "Database access error"
	InvalidRequestFormat = "Invalid request format"
	MissingRequiredField = "Required field not provided"
	InvalidFieldValue    = "Invalid field value"
	RouteNotFound        = "Route not found"
)

// Validation messages
const (
	InvalidEmailFormat = "Invalid email format"
	PasswordTooShort   = "Password must be at least %d characters long"
	UsernameTooShort   = "Username must be at least %d characters long"
	PasswordDigits     = "Password must contain at least digits"
	PasswordUppercase  = "Password must contain at least uppercase letters"
	PasswordLowercase  = "Password must contain at least lowercase letters"
	PasswordSpecial    = "Password must contain at least special characters"
)
