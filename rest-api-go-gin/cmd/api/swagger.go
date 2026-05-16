package main

// GetAllEvents
// @Summary Get all events
// @Description Get all events
// @Tags Events
// @Accept json
// @Produce json
// @Success 200 {array} database.Event
// @Failure 500 {object} gin.H
// @Router /events [get]

// GetEvent
// @Summary Get a specific event
// @Description Get a specific event by ID
// @Tags Events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} database.Event
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /events/{id} [get]

// CreateEvent
// @Summary Create a new event
// @Description Create a new event (requires authentication)
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param event body database.Event true "Event data"
// @Success 201 {object} database.Event
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /events [post]

// UpdateEvent
// @Summary Update an event
// @Description Update an event (requires authentication)
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Event ID"
// @Param event body database.Event true "Event data"
// @Success 200 {object} database.Event
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /events/{id} [put]

// DeleteEvent
// @Summary Delete an event
// @Description Delete an event (requires authentication)
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Event ID"
// @Success 204
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /events/{id} [delete]

// GetAttendeesForEvent
// @Summary Get attendees for an event
// @Description Get all attendees for a specific event
// @Tags Attendees
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {array} database.User
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /events/{id}/attendees [get]

// AddAttendeeToEvent
// @Summary Add attendee to event
// @Description Add an attendee to an event (requires authentication)
// @Tags Attendees
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Event ID"
// @Param userId path int true "User ID"
// @Success 201
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /events/{id}/attendees/{userId} [post]

// DeleteAttendeeFromEvent
// @Summary Remove attendee from event
// @Description Remove an attendee from an event (requires authentication)
// @Tags Attendees
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Event ID"
// @Param userId path int true "User ID"
// @Success 204
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /events/{id}/attendees/{userId} [delete]

// GetEventsByAttendee
// @Summary Get events for an attendee
// @Description Get all events that an attendee is registered for
// @Tags Attendees
// @Accept json
// @Produce json
// @Param id path int true "Attendee ID"
// @Success 200 {array} database.Event
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /attendees/{id}/events [get]

// Login
// @Summary User login
// @Description Login with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body loginRequest true "Login credentials"
// @Success 200 {object} loginResponse
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/login [post]

// RegisterUser
// @Summary User registration
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body registerRequest true "User registration data"
// @Success 201 {object} database.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/register [post]
