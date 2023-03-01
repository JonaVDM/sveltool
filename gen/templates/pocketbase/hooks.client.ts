import { pb } from '$lib/pocketbase'

pb.authStore.loadFromCookie(document.cookie);
