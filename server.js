import { handler } from './build/handler.js';
import https from 'https';
import fs from 'fs';
import express from 'express';

const app = express();

app.use(handler);

let port = 443;
let key = fs.readFileSync('/etc/letsencrypt/live/project.pyroman.io/privkey.pem');
let cert = fs.readFileSync('/etc/letsencrypt/live/project.pyroman.io/fullchain.pem');

https.createServer({
    key: key,
    cert: cert
}, app).listen(port);