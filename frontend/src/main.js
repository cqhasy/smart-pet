import { DialogServer } from '../bindings/changeme/backend/server'

const avatar = document.getElementById('pet-avatar');
const DIALOG_WIDTH = 260;
const DIALOG_HEIGHT = 160;
const GAP = 12;
const SCREEN_PADDING = 20;

function getDialogAnchor() {
    const rect = avatar.getBoundingClientRect();
    const screenX = window.screenX ?? window.screenLeft ?? 0;
    const screenY = window.screenY ?? window.screenTop ?? 0;

    const x = screenX + rect.right + GAP;
    const desiredY = screenY + rect.top - DIALOG_HEIGHT - GAP;
    const y = Math.max(desiredY, screenY + SCREEN_PADDING);

    return { x, y };
}

avatar.addEventListener('click', async () => {
    try {
        const { x, y } = getDialogAnchor();
        await DialogServer.SayHello(x, y);
    } catch (err) {
        console.error('展示对话框失败', err);
    }
});
