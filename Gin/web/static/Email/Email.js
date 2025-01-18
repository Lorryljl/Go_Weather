function createStar() {
    const star = document.createElement('div');
    star.classList.add('star');
    star.style.top = `${Math.random() * 90}%`;
    star.style.left = `${Math.random() * 90}%`;
    document.body.appendChild(star);

    // 每隔一段时间移除星星
    setTimeout(() => {
        document.body.removeChild(star);
    }, 5000); // 星星的运动时间为5秒
}

// 每2秒生成一个五角星
setInterval(createStar, 2000);
