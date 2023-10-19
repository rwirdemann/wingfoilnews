import { writable } from 'svelte/store';
import { browser } from '$app/environment';
let stored = ""
if (browser) {
    stored = localStorage.jwt
}
export const jwt = writable(stored);
if (browser) {
    jwt.subscribe((value) => localStorage.jwt = value)
}