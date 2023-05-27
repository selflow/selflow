describe('Build Workflow', function () {
  Cypress.on('uncaught:exception', (err, runnable) => {
    // returning false here prevents Cypress from
    // failing the test
    return false;
  });

  it('should execute workflow', function () {
    const stepId = 'step-a';
    const stepDockerImage = 'alpine:3.10.0';
    const stepDockerCommands = `echo "##step-a##"\nsleep 5`;

    cy.visit('http://localhost:4200');

    cy.get('#id-input').type(stepId);
    cy.get('#docker-image-input').type(stepDockerImage);
    cy.get('#docker-commands-input').type(stepDockerCommands);

    cy.contains('Create').click();

    cy.contains(stepId).should('be.visible');

    cy.get('#start-run-btn').click();

    cy.get('#full-screen-loader').should('be.visible');

    cy.url().should('contain', '/run');

    cy.contains(stepId, { timeout: 10000 }).should('be.visible').click();

    cy.get('#id-input').should('have.value', stepId);
    cy.get('#docker-image-input').should('have.value', stepDockerImage);
  });
});
