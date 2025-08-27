describe('Full Booking Flow', () => {
  it('should allow a user to login, book a showroom, and see the booking', () => {
    // Visit login page
    cy.visit('/login');
    cy.get('input[type="email"]').type('user@example.com');
    cy.get('input[type="password"]').type('password123');
    cy.get('button[type="submit"]').click();

    // Should redirect to dashboard or showrooms
    cy.url().should('include', '/dashboard');

    // Go to showrooms
    cy.contains('Showrooms').click();
    cy.url().should('include', '/showrooms');

    // Select a showroom (assume first in list)
    cy.get('[data-testid="showroom-card"]').first().click();

    // Select a date in the booking calendar
    cy.get('.react-calendar__tile').not('.booked').first().click();

    // Select a time slot (assume first available)
    cy.get('[data-testid="time-slot"]').not('.booked').first().click();

    // Submit booking
    cy.get('button[type="submit"]').contains(/book/i).click();

    // Check for success message or redirect
    cy.contains(/booking confirmed|my bookings/i);

    // Go to My Bookings
    cy.contains('My Bookings').click();
    cy.url().should('include', '/mybookings');

    // Verify the new booking appears
    cy.contains('user@example.com');
    cy.contains(/showroom/i);
  });
});
