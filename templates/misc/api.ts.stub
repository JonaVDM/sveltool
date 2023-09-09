export class Api {
  private token = '';
  private base = '';

  constructor(url: string) {
    this.base = url;
  }

  public setToken(token: string) {
    this.token = token;
  }

  public async login(email: string, password: string) {
    const req = await this.post('/api/login', { email, password });
    const res = await req.text();

    return res;
  }

  public isLoggedIn(): boolean {
    return this.token != ''
  }

  public toCookie(): string {
    return `auth_token=${this.token}; Max-Age=86400; Path=/`
  }

  public get(url: string) {
    return fetch(this.base + url, {
      headers: {
        'Accept': 'application/json',
        'Authorization': this.token,
      }
    });
  }

  public post(url: string, body: unknown) {
    return fetch(this.base + url, {
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': this.token,
      },
      method: 'POST',
      body: JSON.stringify(body),
    });
  }

  public put(url: string, body: unknown) {
    return fetch(this.base + url, {
      headers: {
        'Accept': 'application/json',
        'Authorization': this.token,
      },
      method: 'PUT',
      body: JSON.stringify(body),
    });
  }

  public delete(url: string) {
    return fetch(this.base + url, {
      headers: {
        'Accept': 'application/json',
        'Authorization': this.token,
      },
      method: 'DELETE',
    });
  }
}

export const api = new Api("http://localhost:8000");
