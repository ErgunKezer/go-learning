fetch('https://www.imperiaonline.org/imperia/game_v6/game/xajax_loader.php', {
    headers: {
        accept: '*/*',
        'accept-language': 'en-US,en;q=0.9,tr;q=0.8',
        'cache-control': 'max-age=0',
        'content-type': 'application/x-www-form-urlencoded',
        'if-modified-since': 'Sat, 1 Jan 2000 00:00:00 GMT',
        'sec-ch-ua':
            '"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
        'sec-fetch-dest': 'empty',
        'sec-fetch-mode': 'cors',
        'sec-fetch-site': 'same-origin',
        cookie: 'PHPSESSID=648b068b294b2e42c7d7b45987a8b20e; iid=328097603; flashless=true; hasflash=false; uiVers=100; lockUiVersion=false; PHPSESSID=9bofdmvao3of9vd1toj942iaci; ref_real=https%253A%252F%252Fwww.imperiaonline.org%252Fimperia%252Fgame_v6%252Fgame%252Fvillage.php; affbid=2755032362; ref=https%3A%2F%2Fwww.imperiaonline.org%2Fimperia%2Fgame_v6%2Fgame%2Fvillage.php',
        Referer:
            'https://www.imperiaonline.org/imperia/game_v6/game/village.php',
        'Referrer-Policy': 'strict-origin-when-cross-origin',
    },
    body: 'xjxfun=viewAlliance&xjxr=1733167539983&xjxargs[]=SallianceInfo&xjxargs[]=%3Cxjxobj%3E%3Ce%3E%3Ck%3Etab%3C%2Fk%3E%3Cv%3EN1%3C%2Fv%3E%3C%2Fe%3E%3Ce%3E%3Ck%3EallianceId%3C%2Fk%3E%3Cv%3EN743%3C%2Fv%3E%3C%2Fe%3E%3Ce%3E%3Ck%3Evexok%3C%2Fk%3E%3Cv%3EBtrue%3C%2Fv%3E%3C%2Fe%3E%3C%2Fxjxobj%3E',
    method: 'POST',
})
    .then((res) => res.text())
    .then((res) => {
        console.log(res);
    });
