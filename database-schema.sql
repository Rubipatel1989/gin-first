-- ============================================================================
-- Championship Coach Platform - Database Schema
-- ============================================================================
-- This file contains the complete database schema for all microservices
-- Database: championship_coach
-- Version: 1.0
-- Date: 2025
-- ============================================================================
-- 
-- PRIMARY KEY STRATEGY:
-- - BIGINT: Tables that can grow to millions+ records (users, events, notifications, etc.)
-- - INT: Tables with medium scale (thousands to hundreds of thousands)
-- - SMALLINT: Small lookup/reference tables (templates, limited categories, organizations)
-- - TINYINT: ENUM replacements (status fields, types, roles, etc.)
-- - Foreign keys use matching type to referenced table
-- ============================================================================

-- Drop database if exists and create fresh
DROP DATABASE IF EXISTS championship_coach;
CREATE DATABASE championship_coach CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE championship_coach;

-- ============================================================================
-- AUTH SERVICE SCHEMA
-- ============================================================================
-- Authentication and Authorization Service
-- ============================================================================

-- Users table (BIGINT - can grow to millions)
CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(50),
    password_hash VARCHAR(255) NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=active, 2=inactive, 3=suspended, 4=pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_email (email),
    INDEX idx_status (status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- User roles (BIGINT - many role assignments)
CREATE TABLE user_roles (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    role TINYINT UNSIGNED NOT NULL COMMENT '1=SuperAdmin, 2=OrgAdmin, 3=Coach, 4=AssistantCoach, 5=Player, 6=Parent',
    organization_id SMALLINT UNSIGNED, -- NULL for SuperAdmin, references organizations
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_role (role),
    INDEX idx_organization_id (organization_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Refresh tokens (BIGINT - many tokens over time)
CREATE TABLE refresh_tokens (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    token VARCHAR(500) NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    revoked_at TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_token (token),
    INDEX idx_expires_at (expires_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Password reset tokens (INT - limited reset attempts)
CREATE TABLE password_resets (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    token VARCHAR(255) NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    used_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_token (token),
    INDEX idx_expires_at (expires_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Player signup requests (for self-registration approval workflow)
CREATE TABLE player_signup_requests (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL, -- User ID (pending player account)
    organization_id SMALLINT UNSIGNED NOT NULL,
    team_id INT UNSIGNED, -- Optional
    requested_by BIGINT UNSIGNED NOT NULL, -- Self (same as user_id)
    reviewed_by BIGINT UNSIGNED, -- Coach/OrgAdmin who reviews
    status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=pending, 2=approved, 3=rejected',
    reviewed_at TIMESTAMP NULL,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_user_id (user_id),
    INDEX idx_organization_id (organization_id),
    INDEX idx_team_id (team_id),
    INDEX idx_status (status),
    INDEX idx_reviewed_by (reviewed_by)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Player invitations (for invitation-based player signup)
CREATE TABLE player_invitations (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    organization_id SMALLINT UNSIGNED NOT NULL,
    team_id INT UNSIGNED, -- Optional
    coach_id BIGINT UNSIGNED NOT NULL, -- User ID of coach
    invitation_code VARCHAR(50) NOT NULL UNIQUE,
    player_email VARCHAR(255),
    player_name VARCHAR(255),
    expires_at TIMESTAMP NULL,
    used_at TIMESTAMP NULL,
    used_by BIGINT UNSIGNED, -- Player user ID if used
    status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=pending, 2=used, 3=expired',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_invitation_code (invitation_code),
    INDEX idx_coach_id (coach_id),
    INDEX idx_organization_id (organization_id),
    INDEX idx_status (status),
    INDEX idx_expires_at (expires_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Parent-Player relationships
CREATE TABLE parent_players (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    parent_id BIGINT UNSIGNED NOT NULL, -- User ID (Parent)
    player_id BIGINT UNSIGNED NOT NULL, -- User ID (Player)
    relationship VARCHAR(50), -- e.g., "Mother", "Father", "Guardian"
    status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=pending, 2=approved, 3=rejected',
    approved_by BIGINT UNSIGNED, -- Player user ID if approved by player
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_parent_id (parent_id),
    INDEX idx_player_id (player_id),
    INDEX idx_status (status),
    UNIQUE KEY unique_parent_player (parent_id, player_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Parent invitations (for invitation-based parent signup)
CREATE TABLE parent_invitations (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    player_id BIGINT UNSIGNED NOT NULL, -- User ID (Player)
    invited_by BIGINT UNSIGNED NOT NULL, -- User ID (Coach or Player)
    invitation_code VARCHAR(50) NOT NULL UNIQUE,
    parent_email VARCHAR(255),
    relationship VARCHAR(50),
    expires_at TIMESTAMP NULL,
    used_at TIMESTAMP NULL,
    used_by BIGINT UNSIGNED, -- Parent user ID if used
    status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=pending, 2=used, 3=expired',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_invitation_code (invitation_code),
    INDEX idx_player_id (player_id),
    INDEX idx_invited_by (invited_by),
    INDEX idx_status (status),
    INDEX idx_expires_at (expires_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- ORGANIZATION SERVICE SCHEMA
-- ============================================================================
-- Organization and Team Management Service
-- ============================================================================

-- Organizations table (SMALLINT - thousands of organizations max, 65,535 limit)
CREATE TABLE organizations (
    id SMALLINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type TINYINT UNSIGNED NOT NULL COMMENT '1=college, 2=club, 3=academy',
    description TEXT,
    logo_url VARCHAR(500),
    address TEXT,
    phone VARCHAR(50),
    email VARCHAR(255),
    website VARCHAR(255),
    status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=active, 2=inactive',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_type (type),
    INDEX idx_status (status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Seasons table (SMALLINT - very limited, maybe 100-200 over many years)
CREATE TABLE seasons (
    id SMALLINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    year INT NOT NULL,
    name VARCHAR(255) NOT NULL, -- e.g., "Fall 2025", "Spring 2026"
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=upcoming, 2=active, 3=completed',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_year (year),
    INDEX idx_status (status),
    INDEX idx_dates (start_date, end_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Teams table (INT - hundreds to thousands per organization)
CREATE TABLE teams (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    organization_id SMALLINT UNSIGNED NOT NULL,
    season_id SMALLINT UNSIGNED NOT NULL,
    name VARCHAR(255) NOT NULL,
    sport_type VARCHAR(100), -- e.g., "Basketball", "Football"
    division VARCHAR(100), -- e.g., "Varsity", "JV", "Freshman"
    description TEXT,
    status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=active, 2=inactive, 3=archived',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE,
    FOREIGN KEY (season_id) REFERENCES seasons(id) ON DELETE RESTRICT,
    INDEX idx_organization_id (organization_id),
    INDEX idx_season_id (season_id),
    INDEX idx_status (status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Coach profiles (INT - one per coach-organization relationship)
CREATE TABLE coach_profiles (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL, -- References users (auth-service)
    organization_id SMALLINT UNSIGNED NOT NULL,
    title VARCHAR(255), -- e.g., "Head Coach", "Assistant Coach"
    external BOOLEAN DEFAULT FALSE, -- True if not on college staff
    status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=active, 2=inactive',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_organization_id (organization_id),
    INDEX idx_status (status),
    UNIQUE KEY unique_coach_org (user_id, organization_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- ROSTER SERVICE SCHEMA
-- ============================================================================
-- Roster and Membership Management Service
-- ============================================================================

-- Team members (BIGINT - many members across all teams)
CREATE TABLE team_members (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    team_id INT UNSIGNED NOT NULL, -- References teams (org-service)
    user_id BIGINT UNSIGNED NOT NULL, -- References users (auth-service)
    member_type TINYINT UNSIGNED NOT NULL COMMENT '1=coach, 2=player',
    jersey_number INT,
    position VARCHAR(100),
    status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=active, 2=inactive, 3=removed, 4=graduated',
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    removed_at TIMESTAMP NULL,
    removed_by BIGINT UNSIGNED, -- User ID who removed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_team_id (team_id),
    INDEX idx_user_id (user_id),
    INDEX idx_status (status),
    INDEX idx_member_type (member_type),
    UNIQUE KEY unique_team_user (team_id, user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Roster lists (INT - limited per team)
CREATE TABLE roster_lists (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    team_id INT UNSIGNED NOT NULL, -- References teams (org-service)
    name VARCHAR(255) NOT NULL,
    purpose TINYINT UNSIGNED NOT NULL COMMENT '1=main, 2=backup, 3=future, 4=tryouts',
    version INT DEFAULT 1,
    created_by BIGINT UNSIGNED NOT NULL, -- User ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_team_id (team_id),
    INDEX idx_purpose (purpose)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Roster list items (BIGINT - many items across all lists)
CREATE TABLE roster_list_items (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    roster_list_id INT UNSIGNED NOT NULL,
    player_id BIGINT UNSIGNED NOT NULL, -- User ID (auth-service)
    priority_rank INT,
    note TEXT,
    added_by BIGINT UNSIGNED NOT NULL, -- User ID
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    removed_at TIMESTAMP NULL,
    FOREIGN KEY (roster_list_id) REFERENCES roster_lists(id) ON DELETE CASCADE,
    INDEX idx_roster_list_id (roster_list_id),
    INDEX idx_player_id (player_id),
    INDEX idx_priority_rank (priority_rank)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Membership history (BIGINT - audit trail can grow large)
CREATE TABLE membership_history (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    team_id INT UNSIGNED NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    action TINYINT UNSIGNED NOT NULL COMMENT '1=added, 2=removed, 3=status_changed',
    old_status VARCHAR(50),
    new_status VARCHAR(50),
    performed_by BIGINT UNSIGNED NOT NULL, -- User ID
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_team_id (team_id),
    INDEX idx_user_id (user_id),
    INDEX idx_action (action),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- EVENT SERVICE SCHEMA
-- ============================================================================
-- Scheduling and Events Service
-- ============================================================================

-- Events table (BIGINT - many events across all teams over time)
CREATE TABLE events (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    team_id INT UNSIGNED NOT NULL, -- References teams (org-service)
    type TINYINT UNSIGNED NOT NULL COMMENT '1=practice, 2=game, 3=meeting, 4=other',
    title VARCHAR(255) NOT NULL,
    description TEXT,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    location VARCHAR(500),
    opponent VARCHAR(255), -- For games
    is_home_game BOOLEAN DEFAULT TRUE, -- For games
    status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=scheduled, 2=cancelled, 3=completed',
    created_by BIGINT UNSIGNED NOT NULL, -- User ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_team_id (team_id),
    INDEX idx_type (type),
    INDEX idx_start_time (start_time),
    INDEX idx_status (status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Game day plans (INT - one per event, but less than events)
CREATE TABLE game_day_plans (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    event_id BIGINT UNSIGNED NOT NULL UNIQUE, -- References events
    reporting_time_jv TIMESTAMP NULL, -- For JV players
    reporting_time_varsity TIMESTAMP NULL, -- For Varsity players
    wear_school JSON, -- Array of checklist items
    wear_game JSON, -- Array of checklist items
    bring_items JSON, -- Array of checklist items
    special_instructions TEXT,
    published_at TIMESTAMP NULL, -- NULL if not published
    published_by BIGINT UNSIGNED, -- User ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
    INDEX idx_event_id (event_id),
    INDEX idx_published_at (published_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Event-announcement linking (INT - junction table)
CREATE TABLE event_announcements (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    event_id BIGINT UNSIGNED NOT NULL, -- References events
    announcement_id BIGINT UNSIGNED NOT NULL, -- References announcements (comm-service)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_event_id (event_id),
    INDEX idx_announcement_id (announcement_id),
    UNIQUE KEY unique_event_announcement (event_id, announcement_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- RSVPs (INT - limited RSVPs per event)
CREATE TABLE event_rsvps (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    event_id BIGINT UNSIGNED NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL, -- References users (auth-service)
    status TINYINT UNSIGNED NOT NULL COMMENT '1=attending, 2=not_attending, 3=maybe',
    notes TEXT,
    responded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
    INDEX idx_event_id (event_id),
    INDEX idx_user_id (user_id),
    UNIQUE KEY unique_event_user (event_id, user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- COMMUNICATION SERVICE SCHEMA
-- ============================================================================
-- Communication and Announcements Service
-- ============================================================================

-- Announcements table (BIGINT - many announcements)
CREATE TABLE announcements (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    team_id INT UNSIGNED NOT NULL, -- References teams (org-service)
    audience_type TINYINT UNSIGNED NOT NULL COMMENT '1=team, 2=group, 3=individual',
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    created_by BIGINT UNSIGNED NOT NULL, -- User ID
    published_at TIMESTAMP NULL, -- NULL if draft
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_team_id (team_id),
    INDEX idx_audience_type (audience_type),
    INDEX idx_published_at (published_at),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Announcement recipients (BIGINT - many recipients)
CREATE TABLE announcement_recipients (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    announcement_id BIGINT UNSIGNED NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL, -- References users (auth-service)
    read_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (announcement_id) REFERENCES announcements(id) ON DELETE CASCADE,
    INDEX idx_announcement_id (announcement_id),
    INDEX idx_user_id (user_id),
    INDEX idx_read_at (read_at),
    UNIQUE KEY unique_announcement_user (announcement_id, user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Announcement attachments (INT - few attachments per announcement)
CREATE TABLE announcement_attachments (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    announcement_id BIGINT UNSIGNED NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_url VARCHAR(500) NOT NULL,
    file_type VARCHAR(100),
    file_size BIGINT, -- Size in bytes
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (announcement_id) REFERENCES announcements(id) ON DELETE CASCADE,
    INDEX idx_announcement_id (announcement_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- NOTIFICATION SERVICE SCHEMA
-- ============================================================================
-- Notification Delivery Service
-- ============================================================================

-- Notifications table (BIGINT - millions of notifications)
CREATE TABLE notifications (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL, -- References users (auth-service)
    type VARCHAR(100) NOT NULL, -- e.g., "GameDayPublished", "Announcement", "ScheduleChange"
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    data JSON, -- Additional data payload
    read_at TIMESTAMP NULL,
    sent_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_user_id (user_id),
    INDEX idx_type (type),
    INDEX idx_read_at (read_at),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Notification templates (SMALLINT - very few templates, maybe 10-20)
CREATE TABLE notification_templates (
    id SMALLINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    type VARCHAR(100) NOT NULL,
    subject VARCHAR(255),
    body_template TEXT NOT NULL, -- Template with placeholders
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_type (type),
    INDEX idx_is_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Notification preferences (INT - one per user)
CREATE TABLE notification_preferences (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL UNIQUE, -- References users (auth-service)
    push_enabled BOOLEAN DEFAULT TRUE,
    email_enabled BOOLEAN DEFAULT FALSE,
    sms_enabled BOOLEAN DEFAULT FALSE,
    game_day_notifications BOOLEAN DEFAULT TRUE,
    schedule_change_notifications BOOLEAN DEFAULT TRUE,
    announcement_notifications BOOLEAN DEFAULT TRUE,
    media_notifications BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_user_id (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- FCM device tokens (INT - few devices per user, maybe 2-5)
CREATE TABLE notification_devices (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL, -- References users (auth-service)
    device_token VARCHAR(500) NOT NULL,
    platform TINYINT UNSIGNED NOT NULL COMMENT '1=ios, 2=android, 3=web',
    device_id VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    last_used_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_user_id (user_id),
    INDEX idx_device_token (device_token),
    INDEX idx_is_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- MEDIA SERVICE SCHEMA
-- ============================================================================
-- Film and Media Management Service
-- ============================================================================

-- Videos table (BIGINT - can grow very large with video uploads)
CREATE TABLE videos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    team_id INT UNSIGNED NOT NULL, -- References teams (org-service)
    event_id BIGINT UNSIGNED, -- References events (event-service), NULL if not linked
    title VARCHAR(255) NOT NULL,
    description TEXT,
    storage_url VARCHAR(500) NOT NULL, -- Cloud Storage URL
    thumbnail_url VARCHAR(500),
    duration INT, -- Duration in seconds
    file_size BIGINT, -- Size in bytes
    file_type VARCHAR(100), -- e.g., "video/mp4"
    processing_status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=pending, 2=processing, 3=completed, 4=failed',
    uploaded_by BIGINT UNSIGNED NOT NULL, -- User ID
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_team_id (team_id),
    INDEX idx_event_id (event_id),
    INDEX idx_processing_status (processing_status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Video tags (BIGINT - many tags across all videos)
CREATE TABLE video_tags (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    video_id BIGINT UNSIGNED NOT NULL,
    player_id BIGINT UNSIGNED NOT NULL, -- User ID (auth-service)
    start_time DECIMAL(10, 2) NOT NULL, -- Start time in seconds
    end_time DECIMAL(10, 2) NOT NULL, -- End time in seconds
    label VARCHAR(255), -- e.g., "Great play", "Needs improvement"
    description TEXT,
    created_by BIGINT UNSIGNED NOT NULL, -- User ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (video_id) REFERENCES videos(id) ON DELETE CASCADE,
    INDEX idx_video_id (video_id),
    INDEX idx_player_id (player_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Video permissions (INT - few permissions per video)
CREATE TABLE video_permissions (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    video_id BIGINT UNSIGNED NOT NULL,
    user_id BIGINT UNSIGNED, -- NULL for public/team access
    permission_type TINYINT UNSIGNED NOT NULL COMMENT '1=public, 2=team, 3=player, 4=parent',
    granted_by BIGINT UNSIGNED NOT NULL, -- User ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (video_id) REFERENCES videos(id) ON DELETE CASCADE,
    INDEX idx_video_id (video_id),
    INDEX idx_user_id (user_id),
    INDEX idx_permission_type (permission_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- STATS SERVICE SCHEMA
-- ============================================================================
-- Statistics and Performance Tracking Service
-- ============================================================================

-- Game statistics (BIGINT - many stat entries)
CREATE TABLE game_stats (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    event_id BIGINT UNSIGNED NOT NULL, -- References events (event-service)
    player_id BIGINT UNSIGNED NOT NULL, -- User ID (auth-service)
    -- Basketball stats (adjust for other sports)
    points INT DEFAULT 0,
    rebounds INT DEFAULT 0,
    assists INT DEFAULT 0,
    steals INT DEFAULT 0,
    blocks INT DEFAULT 0,
    turnovers INT DEFAULT 0,
    fouls INT DEFAULT 0,
    field_goals_made INT DEFAULT 0,
    field_goals_attempted INT DEFAULT 0,
    three_pointers_made INT DEFAULT 0,
    three_pointers_attempted INT DEFAULT 0,
    free_throws_made INT DEFAULT 0,
    free_throws_attempted INT DEFAULT 0,
    minutes_played INT, -- Minutes played
    -- Additional flexible stats (JSON for sport-specific)
    additional_stats JSON,
    notes TEXT,
    entered_by BIGINT UNSIGNED NOT NULL, -- User ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_event_id (event_id),
    INDEX idx_player_id (player_id),
    UNIQUE KEY unique_event_player (event_id, player_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Effort metrics (BIGINT - many metrics)
CREATE TABLE effort_metrics (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    event_id BIGINT UNSIGNED NOT NULL, -- References events (event-service)
    player_id BIGINT UNSIGNED NOT NULL, -- User ID (auth-service)
    hustle_score INT CHECK (hustle_score >= 1 AND hustle_score <= 10),
    engagement_score INT CHECK (engagement_score >= 1 AND engagement_score <= 10),
    notes TEXT,
    entered_by BIGINT UNSIGNED NOT NULL, -- User ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
    INDEX idx_event_id (event_id),
    INDEX idx_player_id (player_id),
    UNIQUE KEY unique_event_player (event_id, player_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- DEVELOPMENT SERVICE SCHEMA
-- ============================================================================
-- Culture and Player Development Service
-- ============================================================================

-- Leadership notes (BIGINT - many notes over time)
CREATE TABLE leadership_notes (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    player_id BIGINT UNSIGNED NOT NULL, -- User ID (auth-service)
    coach_id BIGINT UNSIGNED NOT NULL, -- User ID (auth-service)
    score INT CHECK (score >= 1 AND score <= 10),
    comment TEXT NOT NULL,
    event_id BIGINT UNSIGNED, -- References events (event-service), optional
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_player_id (player_id),
    INDEX idx_coach_id (coach_id),
    INDEX idx_event_id (event_id),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Buy-in scores (BIGINT - many scores)
CREATE TABLE buy_in_scores (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    event_id BIGINT UNSIGNED NOT NULL, -- References events (event-service)
    player_id BIGINT UNSIGNED NOT NULL, -- User ID (auth-service)
    score INT CHECK (score >= 1 AND score <= 10) NOT NULL,
    notes TEXT,
    entered_by BIGINT UNSIGNED NOT NULL, -- User ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_event_id (event_id),
    INDEX idx_player_id (player_id),
    UNIQUE KEY unique_event_player (event_id, player_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Academic entries (INT - few entries per player per term)
CREATE TABLE academic_entries (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    player_id BIGINT UNSIGNED NOT NULL, -- User ID (auth-service)
    term VARCHAR(100) NOT NULL, -- e.g., "Fall 2025", "Semester 1"
    gpa DECIMAL(3, 2), -- GPA score
    notes TEXT,
    entered_by BIGINT UNSIGNED NOT NULL, -- User ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_player_id (player_id),
    INDEX idx_term (term)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Life goals (INT - few goals per player)
CREATE TABLE life_goals (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    player_id BIGINT UNSIGNED NOT NULL, -- User ID (auth-service)
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '1=active, 2=completed, 3=paused',
    target_date DATE,
    completed_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_player_id (player_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- RECRUITING SERVICE SCHEMA
-- ============================================================================
-- Recruiting and Player Profiles Service
-- ============================================================================

-- Recruiting profiles (INT - one per player)
CREATE TABLE recruiting_profiles (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    player_id BIGINT UNSIGNED NOT NULL UNIQUE, -- User ID (auth-service)
    bio TEXT,
    height VARCHAR(50), -- e.g., "6'2\""
    weight INT, -- Weight in pounds
    graduation_year INT,
    position VARCHAR(100),
    additional_info JSON, -- Flexible additional fields
    profile_photo_url VARCHAR(500),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_player_id (player_id),
    INDEX idx_is_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Profile shares (INT - few shares per profile)
CREATE TABLE profile_shares (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    profile_id INT UNSIGNED NOT NULL, -- References recruiting_profiles
    share_token VARCHAR(255) NOT NULL UNIQUE,
    expires_at TIMESTAMP NULL, -- NULL = never expires
    is_active BOOLEAN DEFAULT TRUE,
    view_count INT DEFAULT 0,
    created_by BIGINT UNSIGNED NOT NULL, -- User ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (profile_id) REFERENCES recruiting_profiles(id) ON DELETE CASCADE,
    INDEX idx_share_token (share_token),
    INDEX idx_profile_id (profile_id),
    INDEX idx_is_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Profile permissions (INT - few permissions per profile)
CREATE TABLE profile_permissions (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    profile_id INT UNSIGNED NOT NULL,
    granted_to BIGINT UNSIGNED, -- User ID, NULL for public
    permission_type TINYINT UNSIGNED NOT NULL COMMENT '1=public, 2=coach, 3=college, 4=recruiter',
    granted_by BIGINT UNSIGNED NOT NULL, -- User ID
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (profile_id) REFERENCES recruiting_profiles(id) ON DELETE CASCADE,
    INDEX idx_profile_id (profile_id),
    INDEX idx_granted_to (granted_to)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ============================================================================
-- INITIAL DATA
-- ============================================================================
-- Insert default notification templates
-- ============================================================================

INSERT INTO notification_templates (name, type, subject, body_template, is_active) VALUES
('Game Day Published', 'GameDayPublished', 'Game Day Instructions', 'Game day plan for {{event_title}} has been published. Reporting time: {{reporting_time}}', TRUE),
('New Announcement', 'Announcement', 'New Team Announcement', '{{announcement_title}}\n\n{{announcement_body}}', TRUE),
('Schedule Change', 'ScheduleChange', 'Schedule Change', '{{event_title}} has been {{change_type}}. New time: {{new_time}}', TRUE),
('New Video Tagged', 'VideoTagged', 'New Video Tagged', 'You have been tagged in a new video: {{video_title}}', TRUE);

-- ============================================================================
-- COMMENTS AND NOTES
-- ============================================================================
-- 
-- Database Design Notes:
-- 
-- PRIMARY KEY STRATEGY:
-- 1. BIGINT UNSIGNED: Tables that can grow to millions+ records
--    - users, user_roles, refresh_tokens
--    - team_members, roster_list_items, membership_history
--    - events, announcements, announcement_recipients
--    - notifications, videos, video_tags
--    - game_stats, effort_metrics, leadership_notes, buy_in_scores
-- 
-- 2. INT UNSIGNED: Medium-scale tables (thousands to hundreds of thousands)
--    - teams, coach_profiles
--    - roster_lists, password_resets, game_day_plans
--    - event_announcements, event_rsvps
--    - announcement_attachments, notification_preferences, notification_devices
--    - video_permissions, academic_entries, life_goals
--    - recruiting_profiles, profile_shares, profile_permissions
--    - player_signup_requests, player_invitations, parent_invitations
--    - parent_players (parent-player relationships)
-- 
-- 3. SMALLINT UNSIGNED: Small lookup/reference tables
--    - organizations (max 65,535 organizations)
--    - seasons (limited over time)
--    - notification_templates (very few templates)
-- 
-- 4. TINYINT UNSIGNED: ENUM replacements (status fields, types, roles)
--    All ENUM columns converted to TINYINT with COMMENT showing mapping
--    - users.status: 1=active, 2=inactive, 3=suspended, 4=pending (for player signup requests)
--    - user_roles.role: 1=SuperAdmin, 2=OrgAdmin, 3=Coach, 4=AssistantCoach, 5=Player, 6=Parent
--    - organizations.type: 1=college, 2=club, 3=academy
--    - organizations.status: 1=active, 2=inactive
--    - seasons.status: 1=upcoming, 2=active, 3=completed
--    - teams.status: 1=active, 2=inactive, 3=archived
--    - coach_profiles.status: 1=active, 2=inactive
--    - team_members.member_type: 1=coach, 2=player
--    - team_members.status: 1=active, 2=inactive, 3=removed, 4=graduated
--    - roster_lists.purpose: 1=main, 2=backup, 3=future, 4=tryouts
--    - membership_history.action: 1=added, 2=removed, 3=status_changed
--    - events.type: 1=practice, 2=game, 3=meeting, 4=other
--    - events.status: 1=scheduled, 2=cancelled, 3=completed
--    - event_rsvps.status: 1=attending, 2=not_attending, 3=maybe
--    - announcements.audience_type: 1=team, 2=group, 3=individual
--    - notification_devices.platform: 1=ios, 2=android, 3=web
--    - videos.processing_status: 1=pending, 2=processing, 3=completed, 4=failed
--    - video_permissions.permission_type: 1=public, 2=team, 3=player, 4=parent
--    - life_goals.status: 1=active, 2=completed, 3=paused
--    - profile_permissions.permission_type: 1=public, 2=coach, 3=college, 4=recruiter
--    - player_signup_requests.status: 1=pending, 2=approved, 3=rejected
--    - player_invitations.status: 1=pending, 2=used, 3=expired
--    - parent_invitations.status: 1=pending, 2=used, 3=expired
--    - parent_players.status: 1=pending, 2=approved, 3=rejected
-- 
-- FOREIGN KEY STRATEGY:
-- - Foreign keys use matching integer type to referenced table
-- - organization_id is SMALLINT UNSIGNED (matching organizations.id)
-- - Cross-service references use appropriate integer type (BIGINT for users, SMALLINT for organizations, INT for teams, etc.)
-- - Foreign key constraints are included for data integrity in development
-- - In production with separate databases per service, remove cross-service foreign keys
--   and validate references via API calls instead
-- 
-- STORAGE SAVINGS:
-- - BIGINT: 8 bytes per ID
-- - INT: 4 bytes per ID
-- - SMALLINT: 2 bytes per ID
-- - TINYINT: 1 byte per value (vs ENUM which varies)
-- - Previous UUID approach: 36 bytes per ID
-- - Significant storage savings, especially for high-volume tables
-- 
-- Service Data Isolation:
-- - Each service should only directly access its own tables
-- - Cross-service data access should go through APIs (REST/gRPC)
-- - Foreign key constraints to other services' tables are for development only
-- - For production, split into separate databases and remove cross-service FKs
--
-- Cross-Service References (by ID only, no FK in production):
-- - team_members.team_id → teams (org-service, INT)
-- - team_members.user_id → users (auth-service, BIGINT)
-- - roster_lists.team_id → teams (org-service, INT)
-- - events.team_id → teams (org-service, INT)
-- - announcements.team_id → teams (org-service, INT)
-- - videos.team_id → teams (org-service, INT)
-- - game_stats.event_id → events (event-service, BIGINT)
-- - effort_metrics.event_id → events (event-service, BIGINT)
-- - All user_id references → users (auth-service, BIGINT)
-- - All organization_id references → organizations (org-service, SMALLINT)
--
-- Usage Instructions:
-- 1. For Development: Run this entire script to create all tables in one database
-- 2. For Production: Split into service-specific schemas and remove cross-service FKs
-- 3. Services validate references by calling other services' APIs
-- 4. In application code, use constants/mappings for TINYINT values based on comments
--
-- ============================================================================
