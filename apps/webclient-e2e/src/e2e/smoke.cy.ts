describe('webclient', () => {
  beforeEach(() => cy.visit('/'));

  it('Smoke test', () => {
    expect(1 + 1).to.eq(2);
  });
});
