// CDNを使用する場合のコード
document.addEventListener('DOMContentLoaded', () => {
    // Firebase設定
    const firebaseConfig = {
        apiKey: 'AIzaSyCsjELjHzuSd2j0QYuAPQcCE0nJS4M1HFA',
        authDomain: 'food-shuffle-87dfb.firebaseapp.com',
        projectId: 'food-shuffle-87dfb',
        storageBucket: 'food-shuffle-87dfb.firebasestorage.app',
        messagingSenderId: '966167773179',
        appId: '1:966167773179:web:2a69fffc3cdea4a1ecd55e',
        measurementId: 'G-P8L8M6PTV4',
    };

    // Firebase初期化
    firebase.initializeApp(firebaseConfig);
    const auth = firebase.auth();

    // reCAPTCHAの初期化
    window.recaptchaVerifier = new firebase.auth.RecaptchaVerifier('recaptcha-container', {
        size: 'normal',
        callback: (response) => {
            document.getElementById('register-btn').disabled = false;
        },
    });
    // reCAPTCHA のレンダリング完了時にメッセージを表示
    window.recaptchaVerifier.render();

    const form = document.getElementById('register-form');
    const email = document.getElementById('email');
    const username = document.getElementById('username');
    const password = document.getElementById('password');
    const confirmPassword = document.getElementById('confirm-password');
    const phone = document.getElementById('phone');
    const messageDiv = document.getElementById('message');
    const codeSection = document.getElementById('code-section');
    const codeInput = document.getElementById('code');
    const verifyCodeBtn = document.getElementById('verify-code');
    const infoSection = document.getElementById('info-section');

    let PreRegisterKey = null;

    form.addEventListener('submit', async (e) => {
        e.preventDefault();

        if (password.value !== confirmPassword.value) {
            showMessage('パスワードが一致しません', 'error');
            return;
        }

        try {
            // メール・パスワードでユーザーを作成
            const MailAddress = email.value;
            const UserName = username.value;
            const Password = password.value;
            const ConfirmPassword = confirmPassword.value;
            const Tell = formatPhoneNumber(phone.value);
            console.log(Tell);
            // 仮登録を行う
            const res = await fetch('./v1/pre-register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ MailAddress, UserName, Password, ConfirmPassword, Tell }),
            });

            switch (res.status) {
                case 400:
                    showMessage('項目を正しく入力してください。', 'error');
                    break;
                case 409:
                    showMessage('すでにアカウントが登録されています。', 'error');
                    break;
                case 200:
                    const data = await res.json();
                    console.log(data['Response']['Data']['Key']);

                    PreRegisterKey = data['Response']['Data']['Key'];
                    // 電話番号認証を開始
                    const appVerifier = window.recaptchaVerifier;
                    const confirmationResult = await auth.signInWithPhoneNumber(Tell, appVerifier);

                    window.confirmationResult = confirmationResult;

                    // 認証コード入力画面を表示
                    infoSection.style.display = 'none';
                    codeSection.style.display = 'block';
                    showMessage('認証コードを送信しました', 'success');
                    break;
                case 500:
                    showMessage('サーバー処理中にエラーが発生しました', 'error');
            }
        } catch (error) {
            console.error('Registration error:', error);
            email.value = '';
            username.value = '';
            password.value = '';
            confirmPassword.value = '';
            phone.value = '';
            showMessage(`登録エラー: ${error.message}`, 'error');

            // reCAPTCHAをリセット
            window.recaptchaVerifier.render().then(function (widgetId) {
                grecaptcha.reset(widgetId);
            });
        }
    });

    verifyCodeBtn.addEventListener('click', async () => {
        try {
            const result = await window.confirmationResult.confirm(codeInput.value);

            if (result.user) {
                const Token = await result.user.getIdToken();
                console.log(Token);
                const body = JSON.stringify({ PreRegisterKey, Token });
                console.log(body);
                const res = await fetch('./v1/register', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: body,
                });
                console;
                showMessage('登録が完了しました', 'success');
                const data = await res.json();
                // 登録完了後のリダイレクト
                setTimeout(() => {
                    console.log(`出力するもの:${data['Response']['Data']['JtiToken']}`);
                    JtiTokenChannel.postMessage(data['Response']['Data']['JtiToken']);
                }, 500);
            }
        } catch (error) {
            console.error('Verification error:', error);
            showMessage('認証コードが無効です', 'error');
        }
    });

    function showMessage(message, type) {
        messageDiv.textContent = message;
        messageDiv.className = `message ${type}`;
        messageDiv.style.display = 'block';
    }
});

function formatPhoneNumber(phone) {
    // ハイフンやスペースを除去する
    const cleaned = phone.replace(/[-\s]/g, '');

    // 国際形式の場合（先頭が '+' の場合）
    if (cleaned.startsWith('+')) {
        // +81で始まり、その後9桁または10桁の数字であるかを検証
        if (!/^\+81\d{9,10}$/.test(cleaned)) {
            return 'Invalid phone number';
        }
        return cleaned;
    }

    // 国内の電話番号の場合（先頭が '0' で、全体が10桁または11桁の数字）
    if (!/^0\d{9,10}$/.test(cleaned)) {
        return 'Invalid phone number';
    }

    // 国内番号の場合、先頭の0を除去し+81を付加する
    return '+81' + cleaned.substring(1);
}

