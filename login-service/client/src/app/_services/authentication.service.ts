import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs/Rx';
import 'rxjs/add/operator/map'
import 'rxjs/add/operator/catch'

const httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
@Injectable()
export class AuthenticationService {
    constructor(private http: HttpClient) { }

    apiRoot:string = '/auth';

    login(user) {
        let apiURL = `${this.apiRoot}/login`;
        let body = JSON.stringify(user)
        return this.http.post<any>(apiURL, body, httpOptions)
            .map(user => {
                console.log(user)
                // login successful if there's a jwt token in the response
                if (user && user.token) {
                    // store user details and jwt token in local storage to keep user logged in between page refreshes
                    localStorage.setItem('currentUser', JSON.stringify(user));
                }

                return user;
            })
            .catch(e => {
                console.log(e);
                return Observable.throw(e.error.message);
            });
    }

    logout() {
        // remove user from local storage to log user out
        localStorage.removeItem('currentUser');
    }
}
