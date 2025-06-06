import { Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { CharacterViewerComponent } from './character-viewer/character-viewer.component';
import { TokenViewerComponent } from './token-viewer/token-viewer.component';


export const routes: Routes = [
  { path: '', redirectTo: '/login', pathMatch: 'full' },        // Ruta Raiz, Redirects a Login
  { path: 'login', component: LoginComponent },                 // Login (Generador de Tokens)  
  { path: 'characters', component: CharacterViewerComponent },  // Visor de Personajes
  { path: 'admin', component: TokenViewerComponent },           // Administrador de Tokens
];  