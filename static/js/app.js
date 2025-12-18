const API_BASE = 'http://localhost:8080';

let currentSection = 'users';
let editingId = null;

// Initialize
document.addEventListener('DOMContentLoaded', () => {
    setupNavigation();
    setupModal();
    loadData('users');
});

// Navigation
function setupNavigation() {
    const navItems = document.querySelectorAll('.nav-item');
    navItems.forEach(item => {
        item.addEventListener('click', (e) => {
            e.preventDefault();
            const section = item.dataset.section;
            switchSection(section);
        });
    });
}

function switchSection(section) {
    currentSection = section;
    editingId = null;

    // Update nav
    document.querySelectorAll('.nav-item').forEach(item => {
        item.classList.remove('active');
        if (item.dataset.section === section) {
            item.classList.add('active');
        }
    });

    // Update content
    document.querySelectorAll('.content-section').forEach(sec => {
        sec.classList.remove('active');
    });
    document.getElementById(`${section}-section`).classList.add('active');

    // Update title
    const titles = {
        users: 'Users Management',
        stores: 'Stores Management',
        brands: 'Brands Management'
    };
    document.getElementById('page-title').textContent = titles[section];

    // Load data
    loadData(section);
}

// Load Data
async function loadData(section) {
    const tbody = document.getElementById(`${section}-table-body`);
    tbody.innerHTML = '<tr><td colspan="10" class="loading">Loading...</td></tr>';

    try {
        const response = await fetch(`${API_BASE}/admin/${section}`);
        const data = await response.json();

        if (data.success && data.data) {
            renderTable(section, data.data);
        } else {
            tbody.innerHTML = '<tr><td colspan="10" class="loading">No data found</td></tr>';
        }
    } catch (error) {
        console.error('Error loading data:', error);
        tbody.innerHTML = '<tr><td colspan="10" class="loading">Error loading data</td></tr>';
    }
}

// Render Table
function renderTable(section, items) {
    const tbody = document.getElementById(`${section}-table-body`);
    
    if (items.length === 0) {
        tbody.innerHTML = '<tr><td colspan="10" class="loading">No data found</td></tr>';
        return;
    }

    tbody.innerHTML = items.map(item => {
        if (section === 'users') {
            return `
                <tr>
                    <td>${item.id}</td>
                    <td>${item.name}</td>
                    <td>${item.email}</td>
                    <td>${item.phone || '-'}</td>
                    <td><span class="status-badge ${item.status || 'active'}">${item.status || 'active'}</span></td>
                    <td>
                        <div class="btn-group">
                            <button class="btn btn-primary btn-sm" onclick="editItem('${section}', ${item.id})">Edit</button>
                            <button class="btn btn-danger btn-sm" onclick="deleteItem('${section}', ${item.id})">Delete</button>
                        </div>
                    </td>
                </tr>
            `;
        } else if (section === 'stores') {
            return `
                <tr>
                    <td>${item.id}</td>
                    <td>${item.name}</td>
                    <td>${item.address || '-'}</td>
                    <td>${item.phone || '-'}</td>
                    <td>${item.email || '-'}</td>
                    <td><span class="status-badge ${item.status || 'active'}">${item.status || 'active'}</span></td>
                    <td>
                        <div class="btn-group">
                            <button class="btn btn-primary btn-sm" onclick="editItem('${section}', ${item.id})">Edit</button>
                            <button class="btn btn-danger btn-sm" onclick="deleteItem('${section}', ${item.id})">Delete</button>
                        </div>
                    </td>
                </tr>
            `;
        } else if (section === 'brands') {
            return `
                <tr>
                    <td>${item.id}</td>
                    <td>${item.name}</td>
                    <td>${item.description || '-'}</td>
                    <td>${item.logo ? `<img src="${item.logo}" alt="${item.name}" style="max-width: 50px; max-height: 50px;">` : '-'}</td>
                    <td><span class="status-badge ${item.status || 'active'}">${item.status || 'active'}</span></td>
                    <td>
                        <div class="btn-group">
                            <button class="btn btn-primary btn-sm" onclick="editItem('${section}', ${item.id})">Edit</button>
                            <button class="btn btn-danger btn-sm" onclick="deleteItem('${section}', ${item.id})">Delete</button>
                        </div>
                    </td>
                </tr>
            `;
        }
    }).join('');
}

// Modal
function setupModal() {
    const addBtn = document.getElementById('add-btn');
    const modal = document.getElementById('modal');
    const closeBtn = document.querySelector('.close');
    const cancelBtn = document.getElementById('cancel-btn');
    const form = document.getElementById('modal-form');

    addBtn.addEventListener('click', () => openModal('add'));
    closeBtn.addEventListener('click', closeModal);
    cancelBtn.addEventListener('click', closeModal);
    
    form.addEventListener('submit', handleSubmit);

    window.addEventListener('click', (e) => {
        if (e.target === modal) {
            closeModal();
        }
    });
}

function openModal(mode, item = null) {
    const modal = document.getElementById('modal');
    const title = document.getElementById('modal-title');
    const form = document.getElementById('modal-form');
    
    editingId = mode === 'edit' ? item.id : null;
    title.textContent = mode === 'edit' ? `Edit ${currentSection.slice(0, -1)}` : `Add New ${currentSection.slice(0, -1)}`;

    // Generate form fields based on section
    form.innerHTML = generateFormFields(item);

    modal.classList.add('active');
}

function generateFormFields(item = null) {
    if (currentSection === 'users') {
        return `
            <div class="form-group">
                <label for="name">Name *</label>
                <input type="text" id="name" name="name" value="${item ? item.name : ''}" required>
            </div>
            <div class="form-group">
                <label for="email">Email *</label>
                <input type="email" id="email" name="email" value="${item ? item.email : ''}" required>
            </div>
            <div class="form-group">
                <label for="phone">Phone</label>
                <input type="text" id="phone" name="phone" value="${item ? item.phone : ''}">
            </div>
            ${item ? `
            <div class="form-group">
                <label for="status">Status</label>
                <select id="status" name="status">
                    <option value="active" ${item.status === 'active' ? 'selected' : ''}>Active</option>
                    <option value="inactive" ${item.status === 'inactive' ? 'selected' : ''}>Inactive</option>
                </select>
            </div>
            ` : ''}
        `;
    } else if (currentSection === 'stores') {
        return `
            <div class="form-group">
                <label for="name">Name *</label>
                <input type="text" id="name" name="name" value="${item ? item.name : ''}" required>
            </div>
            <div class="form-group">
                <label for="address">Address</label>
                <input type="text" id="address" name="address" value="${item ? item.address : ''}">
            </div>
            <div class="form-group">
                <label for="phone">Phone</label>
                <input type="text" id="phone" name="phone" value="${item ? item.phone : ''}">
            </div>
            <div class="form-group">
                <label for="email">Email</label>
                <input type="email" id="email" name="email" value="${item ? item.email : ''}">
            </div>
            ${item ? `
            <div class="form-group">
                <label for="status">Status</label>
                <select id="status" name="status">
                    <option value="active" ${item.status === 'active' ? 'selected' : ''}>Active</option>
                    <option value="inactive" ${item.status === 'inactive' ? 'selected' : ''}>Inactive</option>
                </select>
            </div>
            ` : ''}
        `;
    } else if (currentSection === 'brands') {
        return `
            <div class="form-group">
                <label for="name">Name *</label>
                <input type="text" id="name" name="name" value="${item ? item.name : ''}" required>
            </div>
            <div class="form-group">
                <label for="description">Description</label>
                <textarea id="description" name="description">${item ? item.description : ''}</textarea>
            </div>
            <div class="form-group">
                <label for="logo">Logo URL</label>
                <input type="text" id="logo" name="logo" value="${item ? item.logo : ''}" placeholder="https://example.com/logo.png">
            </div>
            ${item ? `
            <div class="form-group">
                <label for="status">Status</label>
                <select id="status" name="status">
                    <option value="active" ${item.status === 'active' ? 'selected' : ''}>Active</option>
                    <option value="inactive" ${item.status === 'inactive' ? 'selected' : ''}>Inactive</option>
                </select>
            </div>
            ` : ''}
        `;
    }
}

function closeModal() {
    const modal = document.getElementById('modal');
    modal.classList.remove('active');
    editingId = null;
}

async function handleSubmit(e) {
    e.preventDefault();
    
    const formData = new FormData(e.target);
    const data = Object.fromEntries(formData.entries());

    try {
        const url = editingId 
            ? `${API_BASE}/admin/${currentSection}/${editingId}`
            : `${API_BASE}/admin/${currentSection}`;
        
        const method = editingId ? 'PUT' : 'POST';

        const response = await fetch(url, {
            method: method,
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
        });

        const result = await response.json();

        if (result.success) {
            closeModal();
            loadData(currentSection);
            alert(result.message || 'Operation successful!');
        } else {
            alert('Error: ' + (result.error || 'Operation failed'));
        }
    } catch (error) {
        console.error('Error:', error);
        alert('Error: ' + error.message);
    }
}

// Edit Item
async function editItem(section, id) {
    try {
        const response = await fetch(`${API_BASE}/admin/${section}/${id}`);
        const result = await response.json();

        if (result.success && result.data) {
            openModal('edit', result.data);
        } else {
            alert('Error loading item');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('Error loading item');
    }
}

// Delete Item
async function deleteItem(section, id) {
    if (!confirm('Are you sure you want to delete this item?')) {
        return;
    }

    try {
        const response = await fetch(`${API_BASE}/admin/${section}/${id}`, {
            method: 'DELETE'
        });

        const result = await response.json();

        if (result.success) {
            loadData(section);
            alert('Item deleted successfully');
        } else {
            alert('Error: ' + (result.error || 'Delete failed'));
        }
    } catch (error) {
        console.error('Error:', error);
        alert('Error: ' + error.message);
    }
}

