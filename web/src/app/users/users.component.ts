import { Component, OnInit } from '@angular/core';
import { User } from '../user/user.component';
import { Config } from '../config';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  users: User[] = [];

  constructor(private title: Title) { }

  ngOnInit() {
    Config.setLogin(false);
    Config.API('users', {}).subscribe(values => this.listUsers(values));
  }

  listUsers(values: any) {
    this.title.setTitle(Config.lang('users') + ' - ' + Config.get('title'));
    Object.entries(values).forEach(user => this.users.push(
      new User(<number>(<unknown>user[0]), <{ [key: string]: Object }>user[1])
    ));
  }
}
