describe('User/Admin Paths & Edge Cases', () => {
  it('should not allow booking the same slot twice (double booking)', () => {
    // Login as user
    cy.visit('/login');
    cy.get('input[type="email"]').type('user@example.com');
    cy.get('input[type="password"]').type('password123');
    cy.get('button[type="submit"]').click();
    cy.url().should('include', '/dashboard');
    cy.contains('Showrooms').click();
    cy.get('[data-testid="showroom-card"]').first().click();
    cy.get('.react-calendar__tile').not('.booked').first().click();
    cy.get('[data-testid="time-slot"]').not('.booked').first().as('slot');
    cy.get('@slot').click();
    cy.get('button[type="submit"]').contains(/book/i).click();
    cy.contains(/booking confirmed|my bookings/i);
    // Try to book the same slot again
    cy.get('@slot').click();
    cy.get('button[type="submit"]').contains(/book/i).click();
    cy.contains(/slot already booked|cannot book/i);
  });

  it('should allow admin to suspend and reactivate a user', () => {
    // Login as admin
    cy.visit('/login');
    cy.get('input[type="email"]').type('admin@example.com');
    cy.get('input[type="password"]').type('adminpass');
    cy.get('button[type="submit"]').click();
    cy.url().should('include', '/dashboard');
    cy.contains('User Manager').click();
    cy.get('[data-testid="user-row"]').first().within(() => {
      cy.contains('Suspend').click();
      cy.contains('Suspended');
      cy.contains('Reactivate').click();
      cy.contains('Active');
    });
  });

  it('should not allow suspended user to book', () => {
    // Login as suspended user
    cy.visit('/login');
    cy.get('input[type="email"]').type('suspended@example.com');
    cy.get('input[type="password"]').type('password123');
    cy.get('button[type="submit"]').click();
    cy.url().should('include', '/dashboard');
    cy.contains('Showrooms').click();
    cy.get('[data-testid="showroom-card"]').first().click();
    cy.get('.react-calendar__tile').not('.booked').first().click();
    cy.get('[data-testid="time-slot"]').not('.booked').first().click();
    cy.get('button[type="submit"]').contains(/book/i).click();
    cy.contains(/account suspended|cannot book/i);
  });
});
