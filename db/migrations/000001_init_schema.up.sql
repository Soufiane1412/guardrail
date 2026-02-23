-- 1. Enbale Row Level Security (RLS)

CREATE TABLE companies (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE,
    domain TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE scans (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    company_id UUID REFERENCES companies(id) ON DELETE CASCADE,
    target_url TEXT NOT NULL,
    status_code INT,
    reputation_score INT CHECK (reputation_score BETWEEN 0 AND 100),
    scan_results JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

--Enable RLS
ALTER TABLE companies ENABLE ROW LEVEL SECURITY;
ALTER TABLE scans ENABLE ROW LEVEL SECURITY;

-- 2. Create a Policy where only 'app_user' can see their own company
CREATE POLICY tenant_isolation_policy ON scans
    USING (company_id = current_setting('app.current_tenant')::uuid);