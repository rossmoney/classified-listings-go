# Full Stack Developer Technical Test - Classified Listings Mini-Application

## Overview
Create a mini classified listings application that demonstrates your full-stack development capabilities. This test is designed to evaluate your skills with Vue.js/Nuxt.js for the frontend and PHP or GoLang for the backend.

## Time Expectation
- We expect this task to take approximately 3-4 hours
- Please don't spend more than 5 hours on this task
- It's okay if you don't complete everything - focus on quality over quantity

## Requirements

### Backend (PHP or GoLang)
Create a RESTful API that provides the following endpoints:

1. GET /api/listings - Retrieve all listings
2. GET /api/listings/{id} - Retrieve a specific listing
3. POST /api/listings - Create a new listing
4. PUT /api/listings/{id} - Update a listing
5. DELETE /api/listings/{id} - Delete a listing

Each listing should have:
- Title
- Description
- Price
- Category (e.g., Property, Vehicle, Electronics)
- Date Posted
- Status (Active/Inactive)

### Frontend (Vue.js/Nuxt.js)
Create a simple but polished interface that includes:

1. A listings page that:
   - Displays all listings in a grid/list view
   - Includes basic filtering by category
   - Implements a simple search by title
   - Shows loading states

2. A create/edit listing form that:
   - Validates input
   - Handles image upload (optional)
   - Provides user feedback on success/error

3. A single listing view that:
   - Shows all listing details
   - Includes a delete confirmation

## Technical Requirements

### Backend
- Implement proper error handling
- Include input validation
- Use appropriate HTTP status codes
- Include basic unit tests
- Document your API (simple README is fine)

### Frontend
- Use Nuxt.js
- Implement proper state management
- Handle loading and error states
- Make it responsive
- Follow Vue.js best practices

### Bonus Points (Optional)
- Docker configuration
- Authentication
- Pagination
- Image upload functionality
- End-to-end tests
- Live search
- Filters persistence

## Submission Guidelines

1. Create a private GitHub repository
2. Include:
   - Complete source code
   - README with:
     - Setup instructions
     - API documentation
     - Any assumptions made
     - What you would do differently in a production environment
   - SQL schema if applicable

## Evaluation Criteria

We will evaluate:

1. Code quality and organization
   - Clean, maintainable code
   - Proper error handling
   - Good commenting and documentation

2. Technical decisions
   - Architecture choices
   - Library/framework usage
   - Database design

3. UI/UX considerations
   - Responsive design
   - Loading states
   - Error handling
   - User feedback

4. Problem solving
   - How you approach the problem
   - Trade-offs made
   - Scalability considerations
