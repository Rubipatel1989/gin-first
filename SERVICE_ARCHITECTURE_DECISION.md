# Service Architecture Decision

**Date:** 2025

## Decision: Separate Admin API and Mobile API Services

We will use **separate microservices** for Admin Panel API and Mobile App API.

## Service Structure

Root directory structure:
- **admin-service** - Admin Panel API (for SuperAdmin/OrgAdmin via Vue.js web portal)
- **mobile-api-service** - Mobile App API (for Coach/Player/Parent via Flutter mobile app)

Both services will use shared packages to avoid code duplication while maintaining independent deployment and scaling.


