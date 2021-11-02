class CustomNetworkError extends Error {
  response: Response

  constructor(name: string, response: Response) {
    super(name)
    this.response = response
  }
}
