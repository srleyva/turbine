import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { User } from '../_models/index';

@Injectable()
export class UserService {
    constructor(private http: HttpClient) { }

    apiRoot:string = '/auth';

    getAll() {
        let apiURL = `${this.apiRoot}/users`;
        return this.http.get<User[]>(apiURL);
    }

    getById(id: number) {
        return this.http.get('/api/users/' + id);
    }

    create(user: User) {
        return this.http.post('/api/users', user);
    }

    update(user: User) {
        return this.http.put('/api/users/' + user.uid, user);
    }

    delete(id: number) {
        return this.http.delete('/api/users/' + id);
    }
}
