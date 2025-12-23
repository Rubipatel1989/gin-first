# Championship Coach Platform - Development Architecture Document

**Version:** 1.0  
**Date:** 2025  
**Audience:** Development Team  
**Status:** Technical Implementation Guide

---

## Table of Contents

1. [Overview](#overview)
2. [Technology Stack](#technology-stack)
3. [Project Structure](#project-structure)
4. [Service Architecture](#service-architecture)
5. [Database Schema](#database-schema)
6. [API Design Standards](#api-design-standards)
7. [Inter-Service Communication](#inter-service-communication)
8. [Authentication & Authorization](#authentication--authorization)
9. [Data Models & Entities](#data-models--entities)
10. [Development Guidelines](#development-guidelines)
11. [Deployment Architecture](#deployment-architecture)
12. [Testing Strategy](#testing-strategy)
13. [Error Handling](#error-handling)
14. [Logging & Monitoring](#logging--monitoring)

---

## Overview

This document provides detailed technical specifications for implementing the Championship Coach Platform. The platform follows a **microservices architecture** pattern where each service is independently deployable and owns its data.

### Architecture Principles

1. **Service Independence**: Each service has its own database and can be developed/deployed independently
2. **API-First**: Clear REST/gRPC contracts between services
3. **Event-Driven**: Asynchronous communication for decoupling
4. **Data Isolation**: Services reference data by ID, not direct database joins
5. **Scalability**: Horizontal scaling per service based on load

---

## Technology Stack

### Backend Services

- **Language**: Go (Golang) 1.21+
- **Web Framework**: Gin (https://github.com/gin-gonic/gin)
- **Database ORM**: GORM (https://gorm.io/)
- **Database**: MySQL 8.0+ (as per business requirements)
- **Caching**: Redis 7.0+
- **Message Broker**: NATS (recommended) or RabbitMQ
- **Authentication**: JWT (JSON Web Tokens)
- **API Documentation**: Swagger/OpenAPI

### Frontend

- **Mobile App**: Flutter (Dart)
- **Admin Portal**: Vue.js 3 + Vite
- **State Management**: Pinia (Vue), Provider/Riverpod (Flutter)

### Infrastructure

- **Containerization**: Docker
- **Orchestration**: Docker Compose (development), Kubernetes (production)
- **Cloud Platform**: Google Cloud Platform (GCP)
- **Object Storage**: Google Cloud Storage (S3-compatible)
- **Push Notifications**: Firebase Cloud Messaging (FCM)

---

## Project Structure

### Repository Organization

```
championship-coach-platform/
├── services/                          # All microservices
│   ├── auth-service/                  # Authentication & Authorization
│   │   ├── cmd/
│   │   │   └── server/
│   │   │       └── main.go
│   │   ├── internal/
│   │   │   ├── handlers/             # HTTP handlers
│   │   │   ├── models/               # Data models
│   │   │   ├── repository/           # Data access layer
│   │   │   ├── service/              # Business logic
│   │   │   ├── middleware/           # Auth middleware
│   │   │   └── config/               # Configuration
│   │   ├── pkg/                      # Shared packages
│   │   ├── database/
│   │   │   ├── mysql.go
│   │   │   └── migrations/
│   │   ├── go.mod
│   │   └── Dockerfile
│   │
│   ├── org-service/                   # Organization & Team Management
│   ├── roster-service/                # Roster & Membership
│   ├── event-service/                 # Scheduling & Events
│   ├── comm-service/                  # Communication
│   ├── notification-service/          # Notifications
│   ├── media-service/                 # Film & Media
│   ├── stats-service/                 # Statistics
│   ├── development-service/           # Culture & Development
│   └── recruiting-service/            # Recruiting & Profiles
│
├── gateway/                           # API Gateway/BFF
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── routes/                   # Route definitions
│   │   ├── middleware/               # Rate limiting, logging
│   │   └── handlers/                 # Gateway handlers
│   └── Dockerfile
│
├── shared/                            # Shared libraries
│   ├── pkg/
│   │   ├── errors/                   # Error definitions
│   │   ├── logger/                   # Logging utilities
│   │   ├── validator/                # Validation utilities
│   │   └── events/                   # Event definitions
│   └── proto/                        # gRPC proto files (if using gRPC)
│
├── scripts/                           # Utility scripts
│   ├── docker-compose.yml
│   ├── setup-db.sh
│   └── run-services.sh
│
├── docs/                              # Documentation
│   ├── api/                          # API documentation
│   └── architecture/                 # Architecture docs
│
└── README.md
```

### Service Structure Template

Each service follows this structure:

```
service-name/
├── cmd/
│   └── server/
│       └── main.go                   # Entry point
├── internal/
│   ├── handlers/                     # HTTP request handlers
│   │   ├── handler.go
│   │   └── routes.go
│   ├── models/                       # Domain models
│   │   └── model.go
│   ├── repository/                   # Data access layer
│   │   └── repository.go
│   ├── service/                      # Business logic
│   │   └── service.go
│   ├── middleware/                   # Service-specific middleware
│   └── config/                       # Configuration
│       └── config.go
├── database/
│   ├── mysql.go                      # Database connection
│   └── migrations/                   # SQL migrations
├── go.mod
├── go.sum
├── .env.example
├── Dockerfile
└── README.md
```

---

## Service Architecture

### 1. Auth Service (`auth-service`)

**Purpose**: Centralized authentication and authorization

**Responsibilities**:
- User authentication (login/logout)
- JWT token generation and validation
- Password management (hashing, reset)
- Role-based access control (RBAC)
- Session management

**Endpoints**:
```
POST   /api/v1/auth/login          # User login
POST   /api/v1/auth/logout         # User logout
POST   /api/v1/auth/refresh        # Refresh access token
POST   /api/v1/auth/forgot-password # Request password reset
POST   /api/v1/auth/reset-password  # Reset password
GET    /api/v1/auth/me             # Get current user
POST   /api/v1/users               # Create user (admin only)
GET    /api/v1/users               # List users (admin only)
GET    /api/v1/users/:id           # Get user (admin only)
PUT    /api/v1/users/:id           # Update user (admin only)
```

**Database Tables**:
- `users` - User accounts
- `user_roles` - User role assignments
- `refresh_tokens` - Refresh token storage
- `password_resets` - Password reset tokens

**Key Features**:
- JWT access tokens (short-lived, 15 minutes)
- Refresh tokens (long-lived, 7 days)
- Password hashing using bcrypt
- Rate limiting on login endpoints

---

### 2. Organization Service (`org-service`)

**Purpose**: Manage organizations, teams, and seasons

**Responsibilities**:
- Organization CRUD operations
- Team management
- Season management
- Coach-organization relationships

**Endpoints**:
```
GET    /api/v1/organizations              # List organizations
POST   /api/v1/organizations              # Create organization
GET    /api/v1/organizations/:id          # Get organization
PUT    /api/v1/organizations/:id          # Update organization
DELETE /api/v1/organizations/:id          # Delete organization

GET    /api/v1/organizations/:id/teams    # List teams in org
POST   /api/v1/organizations/:id/teams    # Create team
GET    /api/v1/teams/:id                  # Get team
PUT    /api/v1/teams/:id                  # Update team
DELETE /api/v1/teams/:id                  # Delete team

GET    /api/v1/seasons                    # List seasons
POST   /api/v1/seasons                    # Create season
GET    /api/v1/seasons/:id                # Get season
PUT    /api/v1/seasons/:id                # Update season

GET    /api/v1/coach-profiles             # List coach profiles
POST   /api/v1/coach-profiles             # Create coach profile
GET    /api/v1/coach-profiles/:id         # Get coach profile
PUT    /api/v1/coach-profiles/:id         # Update coach profile
```

**Database Tables**:
- `organizations` - Organizations (colleges, clubs, academies)
- `teams` - Teams within organizations
- `seasons` - Season definitions
- `coach_profiles` - Coach-organization relationships

---

### 3. Roster Service (`roster-service`)

**Purpose**: Manage team rosters and membership

**Responsibilities**:
- Team member management
- Membership status tracking
- Roster list management
- Membership audit trail

**Endpoints**:
```
GET    /api/v1/teams/:teamId/members              # List team members
POST   /api/v1/teams/:teamId/members              # Add member to team
GET    /api/v1/teams/:teamId/members/:memberId    # Get member
PUT    /api/v1/teams/:teamId/members/:memberId    # Update member status
DELETE /api/v1/teams/:teamId/members/:memberId    # Remove member

GET    /api/v1/teams/:teamId/roster-lists         # List roster lists
POST   /api/v1/teams/:teamId/roster-lists         # Create roster list
GET    /api/v1/roster-lists/:id                   # Get roster list
PUT    /api/v1/roster-lists/:id                   # Update roster list
DELETE /api/v1/roster-lists/:id                   # Delete roster list

GET    /api/v1/roster-lists/:id/items             # List roster items
POST   /api/v1/roster-lists/:id/items             # Add item to roster
PUT    /api/v1/roster-items/:id                   # Update roster item
DELETE /api/v1/roster-items/:id                   # Remove from roster

GET    /api/v1/teams/:teamId/membership-history   # Membership audit trail
```

**Database Tables**:
- `team_members` - Team membership records
- `roster_lists` - Roster list definitions
- `roster_list_items` - Items in roster lists
- `membership_history` - Audit trail of membership changes

**Membership Statuses**:
- `active` - Currently active member
- `inactive` - Temporarily inactive
- `removed` - Removed from team
- `graduated` - Graduated from program

---

### 4. Event Service (`event-service`)

**Purpose**: Event scheduling and game day management

**Responsibilities**:
- Event CRUD operations
- Game day plan management
- Event-announcement linking
- RSVP/attendance tracking (optional)

**Endpoints**:
```
GET    /api/v1/teams/:teamId/events               # List team events
POST   /api/v1/teams/:teamId/events               # Create event
GET    /api/v1/events/:id                         # Get event
PUT    /api/v1/events/:id                         # Update event
DELETE /api/v1/events/:id                         # Delete event

POST   /api/v1/events/:id/game-day                # Create game day plan
GET    /api/v1/events/:id/game-day                # Get game day plan
PUT    /api/v1/events/:id/game-day                # Update game day plan
POST   /api/v1/events/:id/game-day/publish        # Publish game day plan

GET    /api/v1/events/:id/rsvps                   # List RSVPs (if implemented)
POST   /api/v1/events/:id/rsvps                   # Create RSVP
```

**Database Tables**:
- `events` - Event records
- `game_day_plans` - Game day plan details
- `event_announcements` - Link events to announcements (junction table)

**Event Types**:
- `practice` - Team practices
- `game` - Games and matches
- `meeting` - Team meetings
- `other` - Other team events

---

### 5. Communication Service (`comm-service`)

**Purpose**: Team communication and announcements

**Responsibilities**:
- Announcement creation and distribution
- Team-wide, group, or individual messaging
- Announcement targeting

**Endpoints**:
```
GET    /api/v1/teams/:teamId/announcements        # List announcements
POST   /api/v1/teams/:teamId/announcements        # Create announcement
GET    /api/v1/announcements/:id                  # Get announcement
PUT    /api/v1/announcements/:id                  # Update announcement
DELETE /api/v1/announcements/:id                  # Delete announcement

POST   /api/v1/announcements/:id/publish          # Publish announcement
GET    /api/v1/announcements/:id/recipients       # List recipients

GET    /api/v1/users/:userId/announcements        # Get user's announcements
PUT    /api/v1/announcements/:id/read             # Mark as read
```

**Database Tables**:
- `announcements` - Announcement records
- `announcement_recipients` - Recipient tracking
- `announcement_attachments` - File attachments

**Announcement Types**:
- `team` - Team-wide
- `group` - Subgroup (e.g., Varsity only)
- `individual` - Individual message
- `event` - Linked to event

---

### 6. Notification Service (`notification-service`)

**Purpose**: Multi-channel notification delivery

**Responsibilities**:
- Push notification delivery
- Notification template management
- Notification preferences
- Delivery status tracking

**Endpoints**:
```
POST   /api/v1/notifications/send                 # Send notification (internal)
GET    /api/v1/users/:userId/notifications        # Get user notifications
PUT    /api/v1/notifications/:id/read             # Mark as read
DELETE /api/v1/notifications/:id                  # Delete notification

GET    /api/v1/users/:userId/preferences          # Get notification preferences
PUT    /api/v1/users/:userId/preferences          # Update preferences

GET    /api/v1/templates                          # List templates
POST   /api/v1/templates                          # Create template
GET    /api/v1/templates/:id                      # Get template
PUT    /api/v1/templates/:id                      # Update template
```

**Database Tables**:
- `notifications` - Notification records
- `notification_templates` - Notification templates
- `notification_preferences` - User preferences
- `notification_devices` - FCM device tokens

**Event-Driven**:
This service primarily receives events from other services and fans out notifications.

---

### 7. Media Service (`media-service`)

**Purpose**: Video and media management

**Responsibilities**:
- Video upload and storage
- Video metadata management
- Player-specific video tagging
- Video sharing and permissions

**Endpoints**:
```
GET    /api/v1/teams/:teamId/videos               # List team videos
POST   /api/v1/teams/:teamId/videos               # Upload video
GET    /api/v1/videos/:id                         # Get video metadata
PUT    /api/v1/videos/:id                         # Update video
DELETE /api/v1/videos/:id                         # Delete video

POST   /api/v1/videos/:id/tags                   # Add video tag
GET    /api/v1/videos/:id/tags                   # List video tags
DELETE /api/v1/video-tags/:id                    # Delete tag

GET    /api/v1/players/:playerId/videos           # Get player's videos
GET    /api/v1/videos/:id/url                    # Get video URL (signed)
```

**Database Tables**:
- `videos` - Video metadata
- `video_tags` - Player-specific tags/clips
- `video_permissions` - Access control

**Storage**:
- Video files stored in Google Cloud Storage
- Database stores metadata and references

---

### 8. Stats Service (`stats-service`)

**Purpose**: Statistics tracking and analytics

**Responsibilities**:
- Game statistics entry
- Player performance metrics
- Team statistics aggregation
- Effort metrics tracking

**Endpoints**:
```
GET    /api/v1/events/:eventId/stats              # Get event statistics
POST   /api/v1/events/:eventId/stats              # Enter statistics
PUT    /api/v1/stats/:id                          # Update stat
DELETE /api/v1/stats/:id                          # Delete stat

GET    /api/v1/players/:playerId/stats            # Get player stats
GET    /api/v1/teams/:teamId/stats                # Get team stats

POST   /api/v1/events/:eventId/effort             # Enter effort metrics
GET    /api/v1/events/:eventId/effort             # Get effort metrics
GET    /api/v1/players/:playerId/effort           # Get player effort trends
```

**Database Tables**:
- `game_stats` - Game statistics
- `effort_metrics` - Effort tracking (hustle, engagement)
- `stat_definitions` - Custom stat definitions (if needed)

---

### 9. Development Service (`development-service`)

**Purpose**: Track intangible aspects of player development

**Responsibilities**:
- Leadership tracking
- Buy-in tracking
- Academic progress
- Life goals management
- Team mosaic data

**Endpoints**:
```
POST   /api/v1/players/:playerId/leadership       # Add leadership note
GET    /api/v1/players/:playerId/leadership       # Get leadership notes

POST   /api/v1/events/:eventId/buy-in             # Track buy-in
GET    /api/v1/players/:playerId/buy-in           # Get buy-in history

GET    /api/v1/players/:playerId/academics        # Get academic records
POST   /api/v1/players/:playerId/academics        # Add academic entry
PUT    /api/v1/academics/:id                      # Update academic entry

GET    /api/v1/players/:playerId/life-goals       # Get life goals
POST   /api/v1/players/:playerId/life-goals       # Create life goal
PUT    /api/v1/life-goals/:id                     # Update life goal
DELETE /api/v1/life-goals/:id                     # Delete life goal

GET    /api/v1/teams/:teamId/mosaic               # Get team mosaic data
```

**Database Tables**:
- `leadership_notes` - Leadership observations
- `buy_in_scores` - Buy-in tracking
- `academic_entries` - Academic progress
- `life_goals` - Player life goals

---

### 10. Recruiting Service (`recruiting-service`)

**Purpose**: Build comprehensive player profiles for recruiting

**Responsibilities**:
- Profile compilation from multiple services
- Profile sharing and export
- Profile permission management

**Endpoints**:
```
GET    /api/v1/players/:playerId/profile          # Get recruiting profile
PUT    /api/v1/players/:playerId/profile          # Update profile
POST   /api/v1/players/:playerId/profile/share    # Create shareable link
GET    /api/v1/profile-shares/:token              # Get shared profile (public)

GET    /api/v1/players/:playerId/profile/pdf      # Export as PDF
GET    /api/v1/players/:playerId/profile/preview  # Preview profile

GET    /api/v1/profile-permissions                # List permissions
POST   /api/v1/profile-permissions                # Grant permission
DELETE /api/v1/profile-permissions/:id            # Revoke permission
```

**Database Tables**:
- `recruiting_profiles` - Profile metadata
- `profile_shares` - Shareable links
- `profile_permissions` - Access control

**Data Aggregation**:
This service aggregates data from:
- Stats Service (statistics)
- Development Service (leadership, academics)
- Media Service (video clips)

---

## Database Schema

### Database Strategy

**Option 1: Per-Service Database (Recommended for Production)**
- Each service has its own MySQL database
- Complete data isolation
- Independent scaling
- Better fault isolation

**Option 2: Shared Database with Service-Owned Schemas (Development/Early Stage)**
- Single MySQL instance
- Separate schemas per service
- Easier development setup
- Can migrate to per-service databases later

For initial development, we'll use **Option 2** with separate schemas per service.

### Schema Organization

```
Database: championship_coach
├── Schema: auth_db          # Auth Service
├── Schema: org_db           # Organization Service
├── Schema: roster_db        # Roster Service
├── Schema: event_db         # Event Service
├── Schema: comm_db          # Communication Service
├── Schema: notification_db  # Notification Service
├── Schema: media_db         # Media Service
├── Schema: stats_db         # Stats Service
├── Schema: development_db   # Development Service
└── Schema: recruiting_db    # Recruiting Service
```

See `database-schema.sql` for complete schema definitions.

---

## API Design Standards

### RESTful API Conventions

1. **URL Structure**:
   ```
   /api/v1/{resource}
   /api/v1/{resource}/{id}
   /api/v1/{resource}/{id}/{sub-resource}
   ```

2. **HTTP Methods**:
   - `GET` - Read operations
   - `POST` - Create operations
   - `PUT` - Full update operations
   - `PATCH` - Partial update operations
   - `DELETE` - Delete operations

3. **Status Codes**:
   - `200` - Success
   - `201` - Created
   - `204` - No Content (successful delete)
   - `400` - Bad Request
   - `401` - Unauthorized
   - `403` - Forbidden
   - `404` - Not Found
   - `409` - Conflict
   - `422` - Validation Error
   - `500` - Internal Server Error

4. **Request/Response Format**:
   - Content-Type: `application/json`
   - Request body: JSON
   - Response body: JSON

5. **Pagination**:
   ```json
   {
     "data": [...],
     "pagination": {
       "page": 1,
       "limit": 20,
       "total": 100,
       "total_pages": 5
     }
   }
   ```

6. **Error Response Format**:
   ```json
   {
     "error": {
       "code": "VALIDATION_ERROR",
       "message": "Invalid input",
       "details": {
         "email": "Email is required"
       }
     }
   }
   ```

---

## Inter-Service Communication

### Synchronous Communication (REST)

For real-time queries and immediate responses:

**Example**: Event Service needs to validate team exists
```
Event Service → Org Service
GET /api/v1/teams/{teamId}
```

**Implementation**:
- Use HTTP client (net/http or resty)
- Timeout: 5 seconds
- Retry: 3 attempts with exponential backoff
- Circuit breaker pattern for resilience

### Asynchronous Communication (Events)

For workflows and decoupling:

**Event Examples**:
- `GameDayPublished` → Notification Service
- `PlayerAddedToTeam` → Multiple services
- `VideoUploaded` → Media processing queue

**Event Structure**:
```json
{
  "event_type": "GameDayPublished",
  "event_id": "uuid",
  "timestamp": "2025-01-15T10:30:00Z",
  "source_service": "event-service",
  "data": {
    "event_id": "uuid",
    "team_id": "uuid",
    "published_by": "uuid"
  }
}
```

**Message Broker**: NATS (recommended) or RabbitMQ

---

## Authentication & Authorization

### JWT Token Structure

**Access Token**:
```json
{
  "user_id": "uuid",
  "email": "user@example.com",
  "roles": ["coach"],
  "org_ids": ["uuid1", "uuid2"],
  "exp": 1234567890,
  "iat": 1234567890
}
```

**Refresh Token**:
- Stored in database (auth-service)
- Long-lived (7 days)
- Can be revoked

### Authorization Middleware

Each service validates JWT tokens by:
1. Calling Auth Service `/api/v1/auth/validate` (synchronous)
2. OR validating JWT signature locally (if public key available)

**Role-Based Access Control (RBAC)**:
- Roles: `SuperAdmin`, `OrgAdmin`, `Coach`, `AssistantCoach`, `Player`, `Parent`
- Permissions checked at handler level
- Organization-level isolation

---

## Data Models & Entities

### Core Entity Relationships

```
Organization (org-service)
  ├── Teams (org-service)
  │     ├── Team Members (roster-service)
  │     ├── Events (event-service)
  │     ├── Announcements (comm-service)
  │     └── Videos (media-service)
  │
  └── Coach Profiles (org-service)
        └── Users (auth-service)
```

### UUID Strategy

- Use UUID v4 for all primary keys
- Benefits: Distributed generation, no collisions, no sequential IDs

### Soft Deletes

- Use `deleted_at` timestamp (GORM soft delete)
- Allows data recovery and audit trails

### Timestamps

- `created_at` - Record creation time
- `updated_at` - Last update time
- `deleted_at` - Soft delete time (nullable)

---

## Development Guidelines

### Code Organization

1. **Handler Layer**: HTTP request/response handling, validation
2. **Service Layer**: Business logic, orchestration
3. **Repository Layer**: Data access, database operations
4. **Model Layer**: Domain entities, DTOs

### Error Handling

```go
// Define custom error types
type AppError struct {
    Code    string
    Message string
    Status  int
}

// Return errors from service layer
func (s *service) GetTeam(id string) (*Team, error) {
    team, err := s.repo.FindByID(id)
    if err != nil {
        return nil, &AppError{
            Code:    "TEAM_NOT_FOUND",
            Message: "Team not found",
            Status:  http.StatusNotFound,
        }
    }
    return team, nil
}

// Handle in handler layer
func (h *handler) GetTeam(c *gin.Context) {
    team, err := h.service.GetTeam(id)
    if err != nil {
        appErr, ok := err.(*AppError)
        if ok {
            c.JSON(appErr.Status, gin.H{"error": appErr})
            return
        }
        c.JSON(500, gin.H{"error": "Internal server error"})
        return
    }
    c.JSON(200, team)
}
```

### Validation

- Use struct tags for validation (Gin validator)
- Validate at handler layer
- Return detailed validation errors

```go
type CreateEventRequest struct {
    TeamID    string `json:"team_id" binding:"required,uuid"`
    Type      string `json:"type" binding:"required,oneof=practice game meeting other"`
    StartTime string `json:"start_time" binding:"required,datetime"`
    EndTime   string `json:"end_time" binding:"required,datetime"`
}
```

### Database Migrations

- Use GORM AutoMigrate for development
- Use migration files for production
- Version control all migrations

### Environment Configuration

Use environment variables with defaults:

```go
type Config struct {
    ServerPort string `env:"SERVER_PORT" envDefault:"8080"`
    DBHost     string `env:"DB_HOST" envDefault:"localhost"`
    DBPort     string `env:"DB_PORT" envDefault:"3306"`
    DBUser     string `env:"DB_USER" envDefault:"root"`
    DBPassword string `env:"DB_PASSWORD" envDefault:""`
    DBName     string `env:"DB_NAME" envDefault:"championship_coach"`
}
```

---

## Deployment Architecture

### Development Environment

**Docker Compose Setup**:
```yaml
services:
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: championship_coach
    ports:
      - "3306:3306"
  
  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
  
  nats:
    image: nats:latest
    ports:
      - "4222:4222"
  
  auth-service:
    build: ./services/auth-service
    ports:
      - "8080:8080"
    depends_on:
      - mysql
      - redis
  
  # ... other services
```

### Production Environment

- **Kubernetes**: Container orchestration
- **Ingress**: Nginx ingress controller
- **Service Mesh**: Istio (optional)
- **Database**: Cloud SQL (GCP) or managed MySQL
- **Load Balancer**: GCP Load Balancer
- **Auto-scaling**: Horizontal Pod Autoscaling (HPA)

---

## Testing Strategy

### Unit Tests

- Test service layer business logic
- Mock repository layer
- Use `testify` library

### Integration Tests

- Test handler + service + repository
- Use test database
- Clean up after tests

### E2E Tests

- Test complete workflows
- Use test environment
- Mock external services

---

## Error Handling

### Error Types

1. **Validation Errors** (400)
2. **Authentication Errors** (401)
3. **Authorization Errors** (403)
4. **Not Found Errors** (404)
5. **Conflict Errors** (409)
6. **Server Errors** (500)

### Error Logging

- Log all errors with context
- Use structured logging (JSON)
- Include request ID for tracing

---

## Logging & Monitoring

### Structured Logging

```go
log.Info("Event created",
    "event_id", eventID,
    "team_id", teamID,
    "user_id", userID,
)
```

### Log Levels

- `DEBUG` - Detailed information
- `INFO` - General information
- `WARN` - Warning messages
- `ERROR` - Error messages

### Metrics

- Request rate
- Error rate
- Response time (p50, p95, p99)
- Database query time

### Tracing

- Use OpenTelemetry for distributed tracing
- Include trace ID in logs
- Track requests across services

---

## Next Steps

1. **Set up project structure** for all services
2. **Create database schemas** (see `database-schema.sql`)
3. **Implement Auth Service** first (foundation for all other services)
4. **Set up API Gateway** with routing and authentication
5. **Implement remaining services** in priority order
6. **Set up CI/CD pipeline**
7. **Configure monitoring and logging**

---

**Document Status**: Ready for Implementation  
**Questions**: Contact Development Team

