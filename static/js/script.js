document.addEventListener('DOMContentLoaded', function() {
    // Inicializar highlight.js
    hljs.highlightAll();
    
    // Forçar visibilidade de todos os elementos após carregamento
    setTimeout(function() {
        // Forçar visibilidade de todos os elementos
        document.querySelectorAll('section, .endpoint').forEach(function(el) {
            el.style.opacity = '1';
            el.style.transform = 'none';
            el.style.visibility = 'visible';
        });
        
        // Garantir que o conteúdo seja visível
        document.querySelector('.content').style.display = 'block';
    }, 100);

    // Load sidebar
    fetch('static/html/sidebar.html')
        .then(response => response.text())
        .then(data => {
            document.getElementById('sidebar-container').innerHTML = data;
            
            // Adicionar listener para o botão de fechar após carregar a sidebar
            const closeButton = document.getElementById('sidebar-close');
            if (closeButton) {
                closeButton.addEventListener('click', function() {
                    const sidebar = document.querySelector('.sidebar');
                    sidebar.classList.remove('active');
                    document.getElementById('overlay').classList.remove('active');
                });
            }
            
            // Adicionar listeners para os links da sidebar após carregá-la
            document.querySelectorAll('.sidebar a[href^="#"]').forEach(anchor => {
                anchor.addEventListener('click', function (e) {
                    e.preventDefault();
                    
                    const targetId = this.getAttribute('href').substring(1);
                    const targetElement = document.querySelector('#' + targetId);
                    
                    if (targetElement) {
                        // Fechar a sidebar primeiro
                        const sidebar = document.querySelector('.sidebar');
                        sidebar.classList.remove('active');
                        overlay.classList.remove('active');
                        
                        // Atualizar os links ativos
                        updateActiveNavLink(targetId);
                        
                        // Navegação imediata sem atraso
                        const headerOffset = 50;
                        const elementPosition = targetElement.getBoundingClientRect().top;
                        const offsetPosition = elementPosition + window.pageYOffset - headerOffset;
                        
                        window.scrollTo({
                            top: offsetPosition,
                            behavior: 'smooth'
                        });
                    }
                });
            });
        });
    
    // Load footer
    fetch('static/html/footer.html')
        .then(response => response.text())
        .then(data => {
            document.getElementById('footer-container').innerHTML = data;
        })
        .catch(error => console.error('Error loading footer:', error));
    
    // Toggle sidebar
    const sidebarToggle = document.getElementById('sidebar-toggle');
    const overlay = document.getElementById('overlay');
    
    sidebarToggle.addEventListener('click', function() {
        const sidebar = document.querySelector('.sidebar');
        sidebar.classList.toggle('active');
        overlay.classList.toggle('active');
    });
    
    // Close sidebar when overlay is clicked
    overlay.addEventListener('click', function() {
        const sidebar = document.querySelector('.sidebar');
        sidebar.classList.remove('active');
        overlay.classList.remove('active');
    });
    
    // Toggle response containers
    document.addEventListener('click', function(e) {
        if (e.target.closest('.response-toggle')) {
            const toggle = e.target.closest('.response-toggle');
            const content = toggle.nextElementSibling;
            
            toggle.classList.toggle('active');
            content.classList.toggle('active');
        }
    });
    
    // Add glow effect to code blocks
    const codeElements = document.querySelectorAll('pre code');
    codeElements.forEach(el => {
        el.closest('pre').classList.add('code-glow');
    });
    
    // Smooth scroll para links que não estão na sidebar
    document.querySelectorAll('.content a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function (e) {
            e.preventDefault();
            
            const targetId = this.getAttribute('href');
            const targetElement = document.querySelector(targetId);
            
            if (targetElement) {
                window.scrollTo({
                    top: targetElement.offsetTop - 20,
                    behavior: 'smooth'
                });
            }
        });
    });
    
    // Função para atualizar o link ativo na sidebar
    function updateActiveNavLink(targetId = null) {
        // Remover classe active de todos os links
        document.querySelectorAll('.sidebar a').forEach(link => {
            link.classList.remove('active');
        });
        
        if (targetId) {
            // Se um ID foi fornecido (ao clicar), usar este
            const clickedLink = document.querySelector(`.sidebar a[href="#${targetId}"]`);
            if (clickedLink) {
                clickedLink.classList.add('active');
                
                // Se for um subitem, destacar também o item pai
                const parentLi = clickedLink.closest('li').parentNode?.closest('li');
                if (parentLi) {
                    const parentLink = parentLi.querySelector('a');
                    if (parentLink) {
                        parentLink.classList.add('active');
                    }
                }
            }
        }
    }
    
    // Atualizar links ativos no carregamento da página para a seção inicial
    updateActiveNavLink(window.location.hash.substring(1) || 'api-overview');
    
    // Adicionar botão de cópia aos blocos de código
    document.querySelectorAll('pre').forEach(preElement => {
        // Criar botão de cópia
        const copyButton = document.createElement('button');
        copyButton.className = 'copy-button';
        copyButton.innerHTML = '<i class="fas fa-copy"></i> Copy';
        preElement.appendChild(copyButton);
        
        // Adicionar evento de clique para copiar o código
        copyButton.addEventListener('click', function() {
            const codeElement = preElement.querySelector('code');
            const textToCopy = codeElement.textContent;
            
            // Copiar para a área de transferência
            navigator.clipboard.writeText(textToCopy)
                .then(() => {
                    // Mostrar feedback de sucesso
                    copyButton.innerHTML = '<i class="fas fa-check"></i> Copied!';
                    copyButton.classList.add('copied');
                    
                    // Resetar o botão após 2 segundos
                    setTimeout(() => {
                        copyButton.innerHTML = '<i class="fas fa-copy"></i> Copy';
                        copyButton.classList.remove('copied');
                    }, 2000);
                })
                .catch(err => {
                    console.error('Failed to copy text: ', err);
                    copyButton.innerHTML = '<i class="fas fa-times"></i> Error!';
                    
                    // Resetar o botão após 2 segundos
                    setTimeout(() => {
                        copyButton.innerHTML = '<i class="fas fa-copy"></i> Copy';
                    }, 2000);
                });
        });
    });
}); 