document.addEventListener('DOMContentLoaded', function() {
    // Expandir/colapsar seções da barra lateral
    const sectionHeaders = document.querySelectorAll('.section-header');
    
    // Ativar a primeira seção por padrão
    const firstSection = document.querySelector('.nav-section');
    if (firstSection) {
        firstSection.classList.add('active');
    }
    
    sectionHeaders.forEach(header => {
        header.addEventListener('click', function() {
            const section = this.parentElement;
            section.classList.toggle('active');
        });
    });
    
    // Funcionalidade do botão de menu mobile
    const mobileToggle = document.getElementById('mobile-toggle');
    const mobileToggleFixed = document.getElementById('mobile-toggle-fixed');
    const mobileClose = document.getElementById('mobile-close');
    const mobileCloseFixed = document.getElementById('mobile-close-fixed');
    const sidebar = document.querySelector('.sidebar');
    const content = document.querySelector('.content');
    const sidebarOverlay = document.querySelector('.sidebar-overlay');
    const mobileMenuFixed = document.querySelector('.mobile-menu-fixed');
    
    // Função para abrir o menu
    function openMobileMenu() {
        console.log('Opening mobile menu');
        sidebar.classList.add('active');
        if (mobileMenuFixed) {
            mobileMenuFixed.classList.add('active');
        }
        if (sidebarOverlay) {
            sidebarOverlay.style.display = 'block';
        }
        
        // Salvar a posição de rolagem atual
        const scrollY = window.scrollY;
        document.body.style.top = `-${scrollY}px`;
        document.body.classList.add('menu-open');
    }
    
    // Função para fechar o menu
    function closeMobileMenu() {
        console.log('Closing mobile menu');
        sidebar.classList.remove('active');
        if (mobileMenuFixed) {
            mobileMenuFixed.classList.remove('active');
        }
        if (sidebarOverlay) {
            sidebarOverlay.style.display = 'none';
        }
        
        // Restaurar a posição de rolagem
        const scrollY = document.body.style.top;
        document.body.classList.remove('menu-open');
        document.body.style.top = '';
        window.scrollTo(0, parseInt(scrollY || '0') * -1);
    }
    
    // Adicionar eventos aos botões de menu
    if (mobileToggle) {
        console.log('Mobile toggle button found');
        mobileToggle.addEventListener('click', function(e) {
            console.log('Mobile toggle clicked');
            openMobileMenu();
            e.stopPropagation(); // Impedir propagação do evento
        });
    }
    
    if (mobileToggleFixed) {
        console.log('Fixed mobile toggle button found');
        mobileToggleFixed.addEventListener('click', function(e) {
            console.log('Fixed mobile toggle clicked');
            openMobileMenu();
            e.stopPropagation(); // Impedir propagação do evento
        });
    }
    
    if (mobileClose) {
        console.log('Close button found');
        mobileClose.addEventListener('click', function(e) {
            console.log('Close button clicked');
            closeMobileMenu();
            e.stopPropagation(); // Impedir propagação do evento
        });
    }
    
    if (mobileCloseFixed) {
        console.log('Fixed close button found');
        mobileCloseFixed.addEventListener('click', function(e) {
            console.log('Fixed close button clicked');
            closeMobileMenu();
            e.stopPropagation(); // Impedir propagação do evento
        });
    }
    
    // Fechar o menu ao clicar no conteúdo em dispositivos móveis
    if (content) {
        content.addEventListener('click', function() {
            if (window.innerWidth <= 768 && sidebar.classList.contains('active')) {
                closeMobileMenu();
            }
        });
    }
    
    // Fechar o menu ao pressionar a tecla ESC
    document.addEventListener('keydown', function(e) {
        if (e.key === 'Escape' && sidebar.classList.contains('active')) {
            closeMobileMenu();
        }
    });
    
    // Fechar o menu ao clicar em links da navegação em dispositivos móveis
    const mobileNavLinks = document.querySelectorAll('.nav-menu a');
    mobileNavLinks.forEach(link => {
        link.addEventListener('click', function() {
            if (window.innerWidth <= 768 && sidebar.classList.contains('active')) {
                closeMobileMenu();
            }
        });
    });
    
    // Fechar o menu ao clicar no overlay
    if (sidebarOverlay) {
        sidebarOverlay.addEventListener('click', function() {
            console.log('Overlay clicked');
            closeMobileMenu();
        });
    }
    
    // Garantir que o tema escuro esteja sempre ativo
    const htmlElement = document.documentElement;
    htmlElement.classList.add('dark-theme');
    
    // Funcionalidade de troca de idioma
    const languageSelect = document.getElementById('language-select');
    
    // Verificar se há uma preferência de idioma salva
    const savedLanguage = localStorage.getItem('language') || 'en-us';
    htmlElement.setAttribute('lang', savedLanguage);
    
    // Definir o valor do select para o idioma salvo
    if (languageSelect) {
        languageSelect.value = savedLanguage;
        
        // Aplicar as traduções iniciais
        applyTranslations(savedLanguage);
        
        // Adicionar evento de mudança de idioma
        languageSelect.addEventListener('change', function() {
            const selectedLanguage = this.value;
            htmlElement.setAttribute('lang', selectedLanguage);
            localStorage.setItem('language', selectedLanguage);
            
            // Aplicar as traduções para o idioma selecionado
            applyTranslations(selectedLanguage);
        });
    }
    
    // Função para aplicar as traduções
    function applyTranslations(language) {
        const elements = document.querySelectorAll('[data-lang-key]');
        
        elements.forEach(element => {
            const key = element.getAttribute('data-lang-key');
            if (translations[language] && translations[language][key]) {
                // Usar innerHTML em vez de textContent para permitir tags HTML
                element.innerHTML = translations[language][key];
            }
        });
        
        // Atualizar o título da página
        if (translations[language] && translations[language]['api-title']) {
            document.title = translations[language]['api-title'];
        }
        
        // Atualizar os tooltips
        const languageSelector = document.getElementById('language-select');
        if (languageSelector) {
            languageSelector.title = language === 'en-us' ? 'Select language' : 'Selecionar idioma';
        }
    }
    
    // Destacar item ativo na navegação com base na rolagem
    const sections = document.querySelectorAll('.content-section');
    const navLinks = document.querySelectorAll('.subsection a');
    
    function highlightNavOnScroll() {
        let scrollPosition = window.scrollY;
        
        // Adicionar um pequeno offset para melhorar a detecção
        scrollPosition += 100;
        
        sections.forEach(section => {
            const sectionTop = section.offsetTop;
            const sectionHeight = section.offsetHeight;
            const sectionId = section.getAttribute('id');
            
            if (scrollPosition >= sectionTop && scrollPosition < sectionTop + sectionHeight) {
                // Remover classe ativa de todos os links
                navLinks.forEach(link => {
                    link.classList.remove('active');
                });
                
                // Adicionar classe ativa ao link correspondente
                const activeLink = document.querySelector(`.subsection a[href="#${sectionId}"]`);
                if (activeLink) {
                    activeLink.classList.add('active');
                    
                    // Expandir a seção pai se necessário
                    const parentSection = activeLink.closest('.nav-section');
                    if (parentSection && !parentSection.classList.contains('active')) {
                        parentSection.classList.add('active');
                    }
                }
            }
        });
    }
    
    // Adicionar evento de rolagem para destacar navegação
    window.addEventListener('scroll', highlightNavOnScroll);
    
    // Rolagem para links de âncora (sem animação)
    navLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            
            const targetId = this.getAttribute('href').substring(1);
            const targetElement = document.getElementById(targetId);
            
            if (targetElement) {
                // Rolagem instantânea sem comportamento suave
                window.scrollTo(0, targetElement.offsetTop - 20);
                
                // Fechar o menu em dispositivos móveis após clicar
                if (window.innerWidth <= 768 && sidebar.classList.contains('active')) {
                    closeMobileMenu();
                }
                
                // Atualizar a URL com o hash
                history.pushState(null, null, `#${targetId}`);
            }
        });
    });
    
    // Verificar se há um hash na URL ao carregar a página
    if (window.location.hash) {
        const targetId = window.location.hash.substring(1);
        const targetElement = document.getElementById(targetId);
        
        if (targetElement) {
            // Rolagem instantânea sem atraso
            window.scrollTo(0, targetElement.offsetTop - 20);
            
            // Expandir a seção correspondente
            const activeLink = document.querySelector(`.subsection a[href="#${targetId}"]`);
            if (activeLink) {
                const parentSection = activeLink.closest('.nav-section');
                if (parentSection) {
                    parentSection.classList.add('active');
                }
            }
        }
    }

    // Funcionalidade para expandir/colapsar seções de resposta
    function setupResponseSections() {
        // Encontrar todos os elementos h3 com texto "Respostas:" ou "Responses:"
        const responseHeaders = document.querySelectorAll('h3');
        
        responseHeaders.forEach(header => {
            if (header.textContent.trim() === 'Respostas:' || header.textContent.trim() === 'Responses:') {
                // Adicionar atributo data-lang-key
                header.setAttribute('data-lang-key', 'responses');
                
                // Criar a seção de resposta
                const responseSection = document.createElement('div');
                responseSection.className = 'response-section';
                
                // Criar o cabeçalho clicável
                const responseHeaderDiv = document.createElement('div');
                responseHeaderDiv.className = 'response-header';
                
                // Mover o h3 para dentro do cabeçalho
                header.parentNode.insertBefore(responseSection, header);
                responseHeaderDiv.appendChild(header);
                
                // Adicionar o ícone de seta
                const toggleIcon = document.createElement('i');
                toggleIcon.className = 'fas fa-chevron-down response-toggle';
                responseHeaderDiv.appendChild(toggleIcon);
                
                responseSection.appendChild(responseHeaderDiv);
                
                // Criar o contêiner para o conteúdo
                const responseContent = document.createElement('div');
                responseContent.className = 'response-content';
                
                // Mover todo o conteúdo relacionado à resposta para dentro do contêiner
                let nextElement = responseSection.nextSibling;
                
                while (nextElement && 
                      (nextElement.nodeName !== 'H3' || 
                       (nextElement.nodeName === 'H3' && !nextElement.textContent.includes('Respostas:') && !nextElement.textContent.includes('Responses:')))) {
                    
                    // Se encontrarmos outro h3 ou uma nova seção, paramos
                    if ((nextElement.nodeName === 'H3' && nextElement.textContent !== 'Respostas:' && nextElement.textContent !== 'Responses:') || 
                        nextElement.nodeName === 'SECTION') {
                        break;
                    }
                    
                    const elementToMove = nextElement;
                    nextElement = nextElement.nextSibling;
                    responseContent.appendChild(elementToMove);
                }
                
                responseSection.appendChild(responseContent);
                
                // Adicionar evento de clique para expandir/colapsar
                responseHeaderDiv.addEventListener('click', function() {
                    responseSection.classList.toggle('active');
                });
            }
        });
    }
    
    // Executar a configuração das seções de resposta
    setupResponseSections();

    // Configuração do botão do GitHub
    const githubButton = document.getElementById('github-link');
    const GITHUB_REPO = 'https://github.com/pedroborgesdev/hticket-api.git'; // URL do repositório será definida posteriormente

    // Atualizar o link do GitHub quando estiver definido
    if (GITHUB_REPO) {
        githubButton.href = GITHUB_REPO;
        githubButton.style.opacity = '1';
        githubButton.style.pointerEvents = 'auto';
    } else {
        githubButton.style.opacity = '0.5';
        githubButton.style.pointerEvents = 'none';
        githubButton.title = 'GitHub repository not yet defined';
    }
}); 