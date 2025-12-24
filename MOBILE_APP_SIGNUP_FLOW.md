# Mobile App Signup Flow - Players & Parents

**Version:** 1.0  
**Date:** 2025  
**Platform:** Flutter Mobile App  
**Audience:** Players and Parents

---

## Overview

The mobile app is designed for **Coach**, **Player**, and **Parent** users only. SuperAdmin and OrgAdmin use the web admin portal (Vue.js Admin UI).

This document describes the signup and registration flow for **Players** and **Parents** in the mobile app.

---

## 1. Player Signup Flow

### Option A: Self-Registration with Team Selection (Recommended)

```
┌─────────────────────────────────────────────────────────────────────┐
│                    PLAYER SELF-REGISTRATION                          │
└─────────────────────────────────────────────────────────────────────┘

Step 1: Open Mobile App
┌──────────────┐
│ Sign Up      │
│ Button       │
└──────┬───────┘
       │
       ▼

Step 2: Select User Type
┌─────────────────────┐
│ I am a:             │
│                     │
│ [ ] Coach           │
│ [✓] Player          │
│ [ ] Parent          │
└──────┬──────────────┘
       │
       ▼

Step 3: Enter Personal Information
┌─────────────────────┐
│ Sign Up Form        │
│                     │
│ Full Name:          │
│ Email:              │
│ Phone:              │
│ Password:           │
│ Confirm Password:   │
└──────┬──────────────┘
       │
       ▼

Step 4: Search & Select Organization
┌─────────────────────┐
│ Find Your Team      │
│                     │
│ Search Organization:│
│ [ABC College    ]   │
│                     │
│ Results:            │
│ • ABC College       │
│   - Varsity BB      │
│   - JV BB           │
│ • XYZ High School   │
│   - Varsity BB      │
└──────┬──────────────┘
       │
       │ User selects Organization & Team
       ▼

Step 5: Submit Registration
┌─────────────────────┐
│ Submit Request      │
│                     │
│ Request sent to     │
│ coach for approval  │
└──────┬──────────────┘
       │
       │ API Call: POST /api/v1/auth/signup/player
       │ {
       │   name, email, phone, password,
       │   organization_id,
       │   team_id (optional)
       │ }
       ▼

Step 6: Account Created (Pending Status)
┌─────────────────────┐
│ Account Created!    │
│                     │
│ Your request has    │
│ been sent to the    │
│ coach. You will     │
│ receive a           │
│ notification when   │
│ approved.           │
│                     │
│ Status: Pending     │
└─────────────────────┘
       │
       │ Backend creates:
       │ - User account (status: pending)
       │ - User role (role: Player, organization_id)
       │ - Sends notification to coach
       ▼

Step 7: Coach Approval
┌──────────────┐
│ Coach        │
│ (Web/Mobile) │
└──────┬───────┘
       │
       │ Receives notification:
       │ "New player signup request: Jane Player"
       │
       │ Reviews request
       │
       │ Approves/Rejects
       ▼

Step 8: Player Added to Team
┌─────────────────────┐
│ Coach Approves      │
│                     │
│ Backend:            │
│ 1. Update user      │
│    status: active   │
│ 2. Add to team:     │
│    team_members     │
│    (member_type=2)  │
│ 3. Send notification│
│    to player        │
└─────────────────────┘
       │
       │ Player receives notification:
       │ "Welcome! You've been added to Varsity BB"
       ▼

Step 9: Player Can Login
┌─────────────────────┐
│ Player receives     │
│ notification and    │
│ can now login       │
│                     │
│ Login → Access      │
│ granted             │
└─────────────────────┘
```

### Option B: Invitation-Based Registration (Alternative)

```
┌─────────────────────────────────────────────────────────────────────┐
│              PLAYER INVITATION-BASED REGISTRATION                    │
└─────────────────────────────────────────────────────────────────────┘

Step 1: Coach Sends Invitation
┌──────────────┐
│ Coach        │
│ (Web/Mobile) │
└──────┬───────┘
       │
       │ Creates invitation link/code
       │
       │ Shares via:
       │ - Email
       │ - SMS
       │ - Share link
       ▼

Step 2: Player Receives Invitation
┌─────────────────────┐
│ Email/SMS/Link      │
│                     │
│ "You've been        │
│  invited to join    │
│  ABC College        │
│  Varsity BB"        │
│                     │
│ [Accept Invitation] │
└──────┬──────────────┘
       │
       │ Opens mobile app or web link
       ▼

Step 3: Invitation Code Validation
┌─────────────────────┐
│ Enter Invitation    │
│                     │
│ Invitation Code:    │
│ [ABC123-XYZ]        │
│                     │
│ [Validate]          │
└──────┬──────────────┘
       │
       │ API Call: GET /api/v1/invitations/{code}
       │ Returns: organization_id, team_id, coach_name
       ▼

Step 4: Complete Registration
┌─────────────────────┐
│ Sign Up Form        │
│                     │
│ Organization:       │
│ ABC College         │
│ (pre-filled)        │
│                     │
│ Team:               │
│ Varsity BB          │
│ (pre-filled)        │
│                     │
│ Full Name:          │
│ Email:              │
│ Phone:              │
│ Password:           │
│                     │
│ [Sign Up]           │
└──────┬──────────────┘
       │
       │ API Call: POST /api/v1/auth/signup/player
       │ With invitation_code
       ▼

Step 5: Account Created & Auto-Approved
┌─────────────────────┐
│ Welcome!            │
│                     │
│ You've been added   │
│ to Varsity BB       │
│                     │
│ [Go to Dashboard]   │
└─────────────────────┘
       │
       │ Backend:
       │ - Creates user (status: active)
       │ - Creates user_role
       │ - Adds to team_members (auto-approved)
       │ - Marks invitation as used
       ▼

Step 6: Player Can Login
┌─────────────────────┐
│ Player can          │
│ immediately login   │
│ and access app      │
└─────────────────────┘
```

---

## 2. Parent Signup Flow

### Option A: Link to Existing Player Account

```
┌─────────────────────────────────────────────────────────────────────┐
│              PARENT SIGNUP - LINK TO PLAYER                         │
└─────────────────────────────────────────────────────────────────────┘

Step 1: Open Mobile App
┌──────────────┐
│ Sign Up      │
│ Button       │
└──────┬───────┘
       │
       ▼

Step 2: Select User Type
┌─────────────────────┐
│ I am a:             │
│                     │
│ [ ] Coach           │
│ [ ] Player          │
│ [✓] Parent          │
└──────┬──────────────┘
       │
       ▼

Step 3: Link to Player Account
┌─────────────────────┐
│ Link to Player      │
│                     │
│ Option 1:           │
│ Enter Player Email: │
│ [player@email.com]  │
│                     │
│ Option 2:           │
│ Enter Player Code:  │
│ [PLAYER-CODE-123]   │
│                     │
│ [Search/Link]       │
└──────┬──────────────┘
       │
       │ API Call: POST /api/v1/auth/link-player
       │ { player_email or player_code }
       │
       │ Backend validates player exists
       │ Returns: player_name, organization, teams
       ▼

Step 4: Verify Player Information
┌─────────────────────┐
│ Confirm Player      │
│                     │
│ Player:             │
│ Jane Player         │
│                     │
│ Organization:       │
│ ABC College         │
│                     │
│ Team(s):            │
│ Varsity BB          │
│                     │
│ Relationship:       │
│ [Parent/Guardian]   │
│                     │
│ [Confirm]           │
└──────┬──────────────┘
       │
       │ Player receives notification:
       │ "A parent account wants to link to your profile"
       │ Player approves/rejects
       ▼

Step 5: Enter Parent Information
┌─────────────────────┐
│ Parent Information  │
│                     │
│ Full Name:          │
│ Email:              │
│ Phone:              │
│ Password:           │
│ Confirm Password:   │
│                     │
│ [Sign Up]           │
└──────┬──────────────┘
       │
       │ API Call: POST /api/v1/auth/signup/parent
       │ {
       │   name, email, phone, password,
       │   player_id,
       │   relationship
       │ }
       ▼

Step 6: Account Created
┌─────────────────────┐
│ Account Created!    │
│                     │
│ You can now view    │
│ your child's        │
│ information         │
│                     │
│ [Go to Dashboard]   │
└─────────────────────┘
       │
       │ Backend creates:
       │ - User account (role: Parent)
       │ - Links to player via parent_players table
       │ - Sends notification to player
       ▼

Step 7: Parent Can Login
┌─────────────────────┐
│ Parent can login    │
│ and view child's    │
│ data (read-only)    │
└─────────────────────┘
```

### Option B: Invitation from Coach/Player

```
┌─────────────────────────────────────────────────────────────────────┐
│          PARENT SIGNUP - INVITATION FROM COACH/PLAYER               │
└─────────────────────────────────────────────────────────────────────┘

Step 1: Coach/Player Sends Parent Invitation
┌──────────────┐
│ Coach/Player │
└──────┬───────┘
       │
       │ Sends invitation to parent email
       │ "You've been invited to view
       │  Jane Player's profile"
       │
       │ Contains invitation code
       ▼

Step 2: Parent Receives Invitation
┌─────────────────────┐
│ Email               │
│                     │
│ "Coach John has     │
│  invited you to     │
│  view Jane's        │
│  profile"           │
│                     │
│ [Accept Invitation] │
└──────┬──────────────┘
       │
       │ Opens mobile app
       ▼

Step 3: Enter Invitation Code
┌─────────────────────┐
│ Enter Invitation    │
│                     │
│ Invitation Code:    │
│ [PARENT-ABC-XYZ]    │
│                     │
│ [Validate]          │
└──────┬──────────────┘
       │
       │ API Call: GET /api/v1/invitations/{code}
       │ Returns: player_id, player_name, organization
       ▼

Step 4: Complete Parent Registration
┌─────────────────────┐
│ Parent Information  │
│                     │
│ Player:             │
│ Jane Player         │
│ (pre-filled)        │
│                     │
│ Full Name:          │
│ Email:              │
│ Phone:              │
│ Password:           │
│                     │
│ [Sign Up]           │
└──────┬──────────────┘
       │
       │ API Call: POST /api/v1/auth/signup/parent
       │ With invitation_code
       ▼

Step 5: Account Created & Auto-Linked
┌─────────────────────┐
│ Welcome!            │
│                     │
│ You can now view    │
│ Jane's profile      │
│                     │
│ [Go to Dashboard]   │
└─────────────────────┘
       │
       │ Backend:
       │ - Creates user (role: Parent)
       │ - Links to player (auto-approved)
       │ - Marks invitation as used
       ▼

Step 6: Parent Can Login
┌─────────────────────┐
│ Parent can          │
│ immediately login   │
│ and view child's    │
│ data                │
└─────────────────────┘
```

---
## 3. API Endpoints Required

### Player Signup Endpoints:

```
POST   /api/v1/auth/signup/player
  Body: {
    name: string,
    email: string,
    phone: string,
    password: string,
    organization_id: smallint,
    team_id: int (optional),
    invitation_code: string (optional)
  }
  Response: { user_id, status: "pending" or "active" }

GET    /api/v1/organizations/search
  Query: ?q=search_term
  Response: [ { id, name, teams: [...] } ]

POST   /api/v1/invitations/player/validate
  Body: { invitation_code: string }
  Response: { valid: bool, organization_id, team_id, coach_name }

GET    /api/v1/auth/signup-status/{user_id}
  Response: { status: "pending"|"active"|"rejected", message }
```

### Parent Signup Endpoints:

```
POST   /api/v1/auth/signup/parent
  Body: {
    name: string,
    email: string,
    phone: string,
    password: string,
    player_id: bigint (if linking directly),
    player_email: string (if searching),
    invitation_code: string (if using invitation)
  }
  Response: { user_id, parent_id, player_id, status: "pending"|"active" }

POST   /api/v1/auth/link-player
  Body: { player_email: string } OR { player_code: string }
  Response: { player_id, player_name, organization, teams }

POST   /api/v1/invitations/parent/validate
  Body: { invitation_code: string }
  Response: { valid: bool, player_id, player_name, organization }

GET    /api/v1/parents/{parent_id}/requests
  Response: [ { player_id, player_name, status, requested_at } ]

POST   /api/v1/players/{player_id}/approve-parent/{parent_id}
  Response: { success: bool, message }
```

### Coach Endpoints (for managing signups):

```
GET    /api/v1/coaches/{coach_id}/signup-requests
  Response: [ { request_id, user_id, player_name, email, organization, team, requested_at } ]

POST   /api/v1/signup-requests/{request_id}/approve
  Response: { success: bool, user_id, team_member_id }

POST   /api/v1/signup-requests/{request_id}/reject
  Body: { reason: string }
  Response: { success: bool }

POST   /api/v1/teams/{team_id}/invitations/player
  Body: { player_email: string, player_name: string, expires_in_days: int }
  Response: { invitation_code, invitation_link, expires_at }
```

---

## 4. Recommended Approach

### For Players:
**Recommended: Hybrid Approach**
- **Primary**: Self-registration with team selection (Option A)
  - Player searches organization/team
  - Submits request
  - Coach approves and adds to team
- **Secondary**: Invitation-based (Option B)
  - Coach sends invitation
  - Player uses code to signup
  - Auto-approved

### For Parents:
**Recommended: Invitation-Based**
- Coach or Player sends invitation to parent
- Parent uses invitation code to signup
- Auto-linked to player account
- No approval needed (invitation = approval)

---

## 5. User Status Flow

```
┌─────────────────────────────────────────────────────────────────────┐
│                        USER STATUS FLOW                             │
└─────────────────────────────────────────────────────────────────────┘

Player Status Flow:
┌──────────┐
│ Pending  │ ← Player signs up (status=4)
│ (4)      │
└────┬─────┘
     │
     │ Coach approves
     │
     ▼
┌──────────┐     ┌──────────┐
│ Active   │ OR  │ Rejected │
│ (1)      │     │ (3)      │
│          │     └──────────┘
│ Added to │
│ team     │
└──────────┘

Parent Status Flow:
┌──────────┐
│ Pending  │ ← Parent links to player (status=1 in parent_players)
│ (1)      │
└────┬─────┘
     │
     │ Player approves
     │
     ▼
┌──────────┐     ┌──────────┐
│ Approved │ OR  │ Rejected │
│ (2)      │     │ (3)      │
│          │     └──────────┘
│ Can view │
│ player   │
└──────────┘

Note: User account status in users table:
- 1 = active (can login)
- 2 = inactive
- 3 = suspended
- 4 = pending (for players, can't login until approved)
```

---

## 6. Summary

### Player Signup:
1. **Self-Registration**: Player finds organization/team → Submits request → Coach approves → Added to team
2. **Invitation**: Coach sends invitation → Player uses code → Auto-approved → Added to team

### Parent Signup:
1. **Invitation-Based** (Recommended): Coach/Player sends invitation → Parent uses code → Auto-linked → Can view child's data
2. **Link to Player**: Parent searches for player → Requests link → Player approves → Can view child's data

### Key Points:
- ✅ Players can self-register but need coach approval
- ✅ Parents need invitation or player approval
- ✅ Invitation-based signup = auto-approval
- ✅ Self-registration = pending until approval
- ✅ Players join Teams, not Coaches directly
- ✅ Parents link to Players, not Teams

---

**Document Status**: Ready for Implementation  
**Next Steps**: 
1. Update database schema with new tables
2. Implement API endpoints
3. Build mobile app signup UI screens
4. Implement invitation system

